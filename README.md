本项目是以GoFiber为基础的一个Web API开发脚手架，致力于个人快速开发RESTful API

开发日志

- 2021.09.23 Start
- 2021.09.27 Work
- 2021.10.18 Work
- 2021.10.21 Rebuild Project
- 2022.05.04 Open Source

# 已完成

1. 统一返回`JSON`
   1. 200 ok
   2. 302 Redirect
   3. 401 Not Auth
   4. 404 Not Found
   5. 500 Request Failure
   6. 505 Internal Server Error
2. 存储
   1. MySQL Local/Server
   2. Aliyun OSS
   3. Redis Local/Server
   4. Sqlite3 Local File
3. 服务 
   1. Email发送
   2. WebSocket对话
4. 路由分组
   1. v1/v2 API版本控制 
   2. OAuth2 JWT认证控制路由
   3. OAuth2 Github 第三方认证
5. 中间件
   1. 自定义配置Cors
   2. 分布式CSRF   redis-storage
   3. 分布式Session redis-storage
6. 配置
   1. 读取本地配置 `.env`

# 待完成

1. 授权
   1. oauth2 qq  
   2. oauth2 google
   3. oauth2 weibo
2. 认证
   两阶段提交验证
   1. 邮箱验证码
   2. 绑定任意一个第三方平台
   3. 支持短信
3. 存储
   1. etcd
   2. rabbitmq
   3. mysql-gorm
4. 权限模型
   1. acl
   2. rbac
5. 微服务
   1. envoy 
   2. gateway
   3. go-micro
6. 部署运维
   1. 自动伸缩，自动重启
   2. 容器部署，Docker单机,暂不支持K8s
   3. 分布式，暂不支持K8s 
7. 修改工程
   1. 去掉etcd.io dependency
   2. gorm sqlite3 -> mysql 
   3. 增加docker 可运行代码
   4. oauth2_github oauth2_google jwt_auth 等同效果
   5. 发送申请请求，需要认证 注册即登录,header包含 Authorization: bearer {jwt_value}