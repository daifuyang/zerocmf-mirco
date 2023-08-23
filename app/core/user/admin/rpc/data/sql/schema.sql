SET NAMES utf8mb4;
SET
    FOREIGN_KEY_CHECKS = 0;

-- 选择 db:
USE
    zerocmf_user;

-- ----------------------------
-- Table structure for cmf_admin: 管理员基础信息表
-- ----------------------------
# DROP TABLE IF EXISTS `adminUser`;
CREATE TABLE IF NOT EXISTS `admin_user`
(
    `id`             bigint                                                        NOT NULL AUTO_INCREMENT,
    `created_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `status`         tinyint(1)                                                    NOT NULL DEFAULT '1' COMMENT '状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁',
    `desc`           varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述信息',
    `salt`           varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '加密盐',
    `password`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '密码',
    -- 帐号核心信息
    `username`       varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `username_sn`    varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名编号',
    `email`          varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱',
    `mobile_no`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `mobile_country` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '手机号国家码',
    `nick_name`      varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' comment '昵称',
    PRIMARY KEY (`id`),
    UNIQUE idx_username (username),
    KEY `idx_updated_at` (`updated_at`),
    KEY `idx_created_at` (`created_at`)

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='用户信息表';