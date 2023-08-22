---
nav: 指南
group:
  title: 基础
  order: 1
order: 4
---

# 系统函数

默认内置系统函数放在app/std目录下

## 系统初始化

### 函数介绍

```go
func Install(conf database.Mysql, callbacks ...func())
```

### 使用示例

一般放在servicecontext下进行首次运行的初始化操作

```go
import "app/std/util"

util.install(c.Mysql,func() {
	// to do
})
```
