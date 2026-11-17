# zero
## zero service main
## 服务主入口,用于加载微服务组件。

### 插件结构说明
1. 目录与项目结构（以 Domain A 为例）
   每一个领域项目内部结构如下：
```text
domain-a/
├── business/       # 存放业务层相关代码，业务层逻辑，访问数据库等
├── client/         # 领域客户端 SDK（封装了 etcd 发现、负载均衡与 RPC/HTTP 调用）
├── middleware/     # 用于存放中间件，拦截器之类的业务写在这里
├── proto/          # 双方契约（Protobuf 定义），api和rpc
├── rest/           # 领域服务逻辑（负责直接接入服务，网站直接调用，rpc客户端也会服用此逻辑）
├── rpchandles/     # 存放生成rpc通讯逻辑，主要是pb.go和接口
├── server/         # 领域服务端逻辑（包含 etcd 服务注册、业务 handler）
├── types/          # 用于rest或rpc访问的类型定义
└── go.mod
```
* Server 端：负责监听端口、实现业务逻辑、启动时向 etcd 注册自己。
* Client 端：由该领域的开发者写好。其他领域如果想调用 Domain A，不需要重新写调用代码，直接 go get 引入 Domain A 的 client 包即可。Client 内部已经内置好了基于 etcd 的服务发现和负载均衡逻辑。

按规范编写的微服务插件，通过import 领域微服务模块+配置文件配置，即可完成加载。
自动生成的代码类，一般是类名全小写，我们这里不作修改，按小写拷贝放入。
自己手动建的类，采用驼峰命名法，并且首字母大写。

本项目的必须组件：
```text
zero-user   用户与权限管理
zero-web-core web框架内核
zero-core 框架内核
```

本项目的必须组件：
```text
zero-file 文件管理，上传、下载、存取区域、版本管理等
```