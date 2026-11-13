# zero
## zero service main
## 服务主入口,用于加载微服务组件。

### 插件结构说明
1. 目录与项目结构（以 Domain A 为例）
   每一个领域项目内部结构如下：
```text
domain-a/
├── client/          # 领域客户端 SDK（封装了 etcd 发现、负载均衡与 RPC/HTTP 调用）
├── domain/          # 领域服务逻辑（不适用etcd时，直接调用）
├── server/          # 领域服务端逻辑（包含 etcd 服务注册、业务 handler）
├── proto/           # 双方契约（Protobuf 定义）
└── go.mod
```
* domain : 负责直接接入服务。
* Server 端：负责监听端口、实现业务逻辑、启动时向 etcd 注册自己。
* Client 端：由该领域的开发者写好。其他领域如果想调用 Domain A，不需要重新写调用代码，直接 go get 引入 Domain A 的 client 包即可。Client 内部已经内置好了基于 etcd 的服务发现和负载均衡逻辑。