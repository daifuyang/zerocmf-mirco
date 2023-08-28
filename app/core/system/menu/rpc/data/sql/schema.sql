SET NAMES utf8mb4;
SET
    FOREIGN_KEY_CHECKS = 0;

-- 选择 db:
USE
    zerocmf_system;

-- ----------------------------
-- Table structure for menu: 系统菜单
-- ----------------------------
# DROP TABLE IF EXISTS `admin_menu`;
create table IF NOT EXISTS admin_menu
(
    `id`           bigint                                                        NOT NULL AUTO_INCREMENT,
    `parent_id`    bigint                                                        NOT NULL DEFAULT 0 COMMENT '上级id',
    `user_id`      bigint                                                        NOT NULL DEFAULT 0 COMMENT '创建人',
    `form_id`      varchar(100)                                                  NOT NULL DEFAULT 0 COMMENT '表单id',
    `created_at`   datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`   datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `menu_type`    tinyint(2)                                                    NOT NULL DEFAULT 0 COMMENT '菜单类型（0：分组，1：菜单，2：表单，3：按钮）',
    `name`         varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '菜单名称',
    `icon`         varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '菜单图标',
    `path`         varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路由地址',
    `component`    varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '组件地址',
    `access`       varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '权限字符',
    `link`         tinyint(2)                                                    NOT NULL DEFAULT 0 COMMENT '是否外链：0：否，1：是',
    `order`        float(0)                                                      NOT NULL DEFAULT 10000 COMMENT '排序，越大越靠前',
    `hide_in_menu` tinyint(2)                                                    NOT NULL DEFAULT 0 COMMENT '菜单中隐藏',
    `status`       tinyint(2)                                                    NOT NULL DEFAULT 1 COMMENT '1 =>启用,0 => 停用',
    PRIMARY KEY (`id`)

)
    comment '系统菜单';