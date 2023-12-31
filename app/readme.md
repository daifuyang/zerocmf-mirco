## 服务编码:

- ✅ 服务的端口号 = 服务编号 + 服务类型编码

| 服务类型编码 | 说明        | 重要性  |
|--------|-----------|------|
| 0      | rpc       | ⭐⭐⭐️ |
| 1      | job       | ⭐    |
| 2      | admin api | ⭐    |
| 8      | web api   | ⭐⭐⭐  | 

## 服务列表:

| 服务                    | 编号   | 说明   | 依赖服务 | 服务等级 |
|-----------------------|------|------|------|------|
| [user](core/user)     | 101X | 用户中心 | 无    | L0   |
| [system](core/system) | 102X | 系统中心 | 无    | L0   |

## api服务：

| 服务                                    | 编号    | 说明   | 依赖服务 | 服务等级 |
|---------------------------------------|-------|------|------|------|
| [user/admin/api](core/user/admin/api) | 10108 | 后台用户 | 无    | L0   |
| [user/authn/api](core/user/authn/api) | 10118 | 身份验证 | 无    | L0   |

## rpc服务

| 服务                                      | 编号    | 说明   | 依赖服务 | 服务等级 |
|-----------------------------------------|-------|------|------|------|
| [user/admin/rpc](core/user/admin/rpc)   | 10100 | 后台用户 | 无    | L0   |
| [system/log/rpc](core/system/log/rpc)   | 10200 | 身份验证 | 无    | L0   |
| [system/menu/rpc](core/system/menu/rpc) | 10210 | 菜单管理 | 无    | L0   |

## 服务端口:

> eg: [user/admin](core/user/admin)

| 服务类型      | 端口号              | 说明           |
|-----------|------------------|--------------|
| rpc       | 10100 = 1010 + 0 | rpc 服务       |
| job       | 10101 = 1010 + 1 | job 服务       |
| admin api | 10102 = 1010 + 2 | admin api 服务 |
| user api  | 10108 = 1010 + 8 | user api 服务  |

