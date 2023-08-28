SET NAMES utf8mb4;
SET
    FOREIGN_KEY_CHECKS = 0;

-- 选择 db:
USE
    zerocmf_system;

-- ----------------------------
-- Table structure for system_log: 系统运行日志
-- ----------------------------
# DROP TABLE IF EXISTS `system_log`;
create table IF NOT EXISTS system_log
(
    `id`         bigint                                                        NOT NULL AUTO_INCREMENT,
    `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `app_name`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '应用名称',
    `logLevel`   varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '日志级别，如 INFO、ERROR、WARNING 等',
    `method`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '日志发生所在方法',
    `message`    text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci         NOT NULL DEFAULT '' COMMENT '日志消息',
    PRIMARY KEY (`id`)

)
    comment '系统运行日志';

-- ----------------------------
-- Table structure for login_log: 管理员登录日志表
-- ----------------------------
create table IF NOT EXISTS login_log
(
    `id`         bigint                                                        NOT NULL AUTO_INCREMENT,
    `created_at` datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `user_id`    bigint                                                        NOT NULL DEFAULT 0 COMMENT '用户id',
    `username`   varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `status`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' comment '登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）',
    `ip`         varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' comment 'IP地址',
    PRIMARY KEY (`id`)
)
    comment '系统登录日志';

-- ----------------------------
-- Table structure for operation_log: 管理员操作日志表
-- ----------------------------
# DROP TABLE IF EXISTS `operation_log`;
create table IF NOT EXISTS operation_log
(
    `id`              bigint                                                         NOT NULL AUTO_INCREMENT,
    `operation_at`    datetime                                                       NOT NULL DEFAULT CURRENT_TIMESTAMP comment '操作时间',
    `user_id`         varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '用户id',
    `username`        varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '用户名',
    `operation`       varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci   NOT NULL DEFAULT '' COMMENT '用户操作',
    `method`          varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '请求方法',
    `request_params`  varchar(5000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '请求参数',
    `response_params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci          NOT NULL DEFAULT '' COMMENT '响应参数',
    `time`            bigint                                                         NOT NULL DEFAULT 0 COMMENT '执行时长(毫秒)',
    `ip`              varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci   NOT NULL DEFAULT '' COMMENT 'IP地址',
    PRIMARY KEY (`id`),
    KEY `idx_operation_at` (`operation_at`)
)
    comment '系统操作日志';