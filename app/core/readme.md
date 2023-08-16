# core:

- 基础公共服务: 通用服务, 如 用户注册/登录/鉴权等, 短信, 邮件, 推送, 通知等

## 服务编码:

- ✅ 服务的端口号 = 服务编号 + 服务类型编码

| 服务类型编码 | 说明        | 重要性  |
|--------|-----------|------|
| 0      | rpc       | ⭐⭐⭐️ |
| 1      | job       | ⭐    |
| 2      | admin api | ⭐    |
| 8      | web api   | ⭐⭐⭐  | 

## 服务列表:

| 服务                           | 编号   | 说明     | 依赖服务 | 服务等级 |
|------------------------------|------|--------|------|------|
| [user](user)                 | 101X | 用户中心   | 无    | L0   |
| [notification](notification) | 102X | 消息推送中心 | 无    | L0   |
| [security](security)         | 103X | 安全中心   | 无    | L0   |
| [id-allocator](id-allocator) | 104X | 发号器    | 无    | L0   |

## 服务端口:

> eg: [user/member](user/member)

| 服务类型      | 端口号              | 说明           |
|-----------|------------------|--------------|
| rpc       | 10100 = 1010 + 0 | rpc 服务       |
| job       | 10101 = 1010 + 1 | job 服务       |
| admin api | 10102 = 1010 + 2 | admin api 服务 |
| user api  | 10108 = 1010 + 8 | user api 服务  |