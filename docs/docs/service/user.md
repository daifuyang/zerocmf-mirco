---
nav:
  title: 服务模块
  order: 1
group:
  title: 核心模块
  order: 0
order: 1
---

# 用户模块

## 身份验证模块

服务类型：api  
服务端口：10118  
服务介绍：身份验证相关功能

| 注册路由 | 功能描述 | 备注 |
|---------------------------|-------|--|
| /api/v1/authn/admin/login | 管理员登录 | |
| /api/v1/authn/validation | 身份验证 | 配合apisix的forward插件进行身份验证 |

## 管理员模块
