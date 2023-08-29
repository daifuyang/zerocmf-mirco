package util

import (
	"app/std/database"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
	"io"
	"os"
	"strings"
)

// 首次执行初始化程序

func Install(callbacks func(), debug ...bool) {

	folderPath := "data"
	_, folderErr := os.Stat(folderPath)
	if os.IsNotExist(folderErr) {
		// 文件夹不存在，创建文件夹
		folderErr = os.Mkdir(folderPath, 0755) // 0755 是文件夹的权限
		if folderErr != nil {
			fmt.Println("无法创建文件夹:", folderErr)
			return
		}
		fmt.Println("文件夹已创建:", folderPath)
	}

	filePath := "data/install"
	// 检查文件是否存在
	_, fileNotExistErr := os.Stat(filePath)
	if os.IsNotExist(fileNotExistErr) {

		callbacks()

		if debug == nil {
			// 文件不存在，创建文件
			file, createErr := os.Create(filePath)
			if createErr != nil {
				fmt.Println("无法创建文件:", createErr)
				return
			}
			err := file.Close()
			if err != nil {
				fmt.Println("file close err:", err.Error())
				return
			}
		}
		fmt.Println("初始化成功！")
	} else if fileNotExistErr == nil {
		fmt.Println("服务已安装！")
	} else {
		fmt.Println("发生错误:", fileNotExistErr)
	}
}

func InitDb(conf database.Mysql) {
	charset := conf.Charset
	if strings.TrimSpace(charset) == "" {
		charset = "utf8mb4"
	}
	collate := conf.Collate
	if strings.TrimSpace(collate) == "" {
		collate = "utf8mb4_unicode_ci"
	}
	if strings.TrimSpace(conf.DbName) == "" {
		return
	}
	// 执行初始化数据库任务
	createTable := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET `%s` DEFAULT COLLATE `%s`;", conf.DbName, charset, collate)
	dsn := conf.Dsn(conf.WithSimple())
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("mysql数据库异常：" + err.Error())
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println("关闭数据库连接出错:", err.Error())
		}
	}(db)
	_, err = db.Exec(createTable)
	if err != nil {
		panic("建表失败：" + err.Error())
	}

	// 执行数据库初始化迁移
	sqlPath := "data/sql/schema.sql"
	file, fileOpenErr := os.Open(sqlPath)
	if fileOpenErr != nil {
		fmt.Println("跳过初初始化数据库！")
	} else {
		defer func(file *os.File) {
			fileErr := file.Close()
			if fileErr != nil {
				fmt.Println("file close err:", fileErr.Error())
			}
		}(file)
		content, readErr := io.ReadAll(file)
		if readErr != nil {
			panic("读取文件失败:" + readErr.Error())
		}
		statements := strings.Split(string(content), ";")
		for _, statement := range statements {
			cleanedStatement := strings.TrimSpace(statement)
			if cleanedStatement != "" {
				_, execErr := db.Exec(cleanedStatement)
				if execErr != nil {
					fmt.Println("cleanedStatement", cleanedStatement)
					panic("执行数据库失败:" + execErr.Error())
				}
			}
		}
	}
}
