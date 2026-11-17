# zero
## zero service main
## 服务主入口,用于加载微服务组件。

### 领域业务插件结构说明
1. 目录与项目结构（以 Domain A 为例）
   每一个领域项目内部结构如下：
```text
domain-a/
├── business/       # 存放业务层相关代码，业务层逻辑，访问数据库等
├── client/         # 领域客户端 SDK（封装了 etcd 发现、负载均衡与 RPC/HTTP 调用）
├──────── rpc/      # 存放通过proto生成的client/xxxxservice.go，为了方便引用，我们将父文件夹设置成xxxxclient/
├── middleware/     # 用于存放中间件，拦截器之类的业务写在这里
├── proto/          # 双方契约（Protobuf 定义），api和rpc
├──────── api/      # 用于生成rest的api定义文件
├──────── rpc/      # 用于生成rpc服务的服务定义文件，主要是双方契约（Protobuf 定义）    
├── rest/           # 领域服务逻辑（负责直接接入服务，网站直接调用，rpc客户端也会服用此逻辑）
├──────── action/   # 存放通过api生成的logic目录内的内容，这里已经将同领域的方法合并到一个文件内。
├──────── handler/  # 存放通过api生成的routes.go，并把所有的需要调用的handler已经合并到一个文件内。
├── rpchandles/     # 存放生成rpc通讯逻辑，主要是pb.go和接口
├── server/         # 领域服务端逻辑（包含 etcd 服务注册、业务 handler）
├──────── action/   # 存放通过proto生成的logic目录内的内容，这里已经将同领域的方法合并到一个文件内。
├──────── rpc/      # 存放通过proto生成的server/xxxxservice.go,为了方便引用，我们将父文件夹设置成xxxxserver/
├── types/          # 用于rest或rpc访问的类型定义，需要删除一些公共类型，以免冗余。
└── go.mod
```
* Rest 端：必须，实现网站接入逻辑，如果不提供单体网站服务，在action中不要对业务层逻辑进行接入，直接调用rpc client接入。
* Server 端：可选，负责监听端口、实现业务逻辑、启动时向 etcd（也可以是其他注册中心） 注册自己。
* Client 端：可选，复用Rest端逻辑，在action中通过判断服务类型是RpcClient进行rpc client接入，接入后启动时会向etcd 请求对应的服务。

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