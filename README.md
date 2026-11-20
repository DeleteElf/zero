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
├──────── routes/   # 存放通过api生成的routes.go，并把所有的需要调用的handler已经合并到一个文件内。
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



#### 参考：
1. https://blog.csdn.net/weixin_36865481/article/details/140481869
2. https://blog.csdn.net/weixin_46453070/article/details/141461552

初始化项目

go mod init namespace

配置代理
```shell
go env -w GOPROXY=https://goproxy.cn,direct
```

清除代理
```shell
go env -u GOPROXY
```

# 私有仓库不走 GOPROXY，直连 GitHub
go env -w GOPRIVATE=github.com/*

# 或者指定具体组织
go env -w GOPRIVATE=cgtw.com/service/*,github.com/DeleteElf/*

# 允许http拉取
go env -w GOINSECURE=cgtw.com/*,github.com/DeleteElf/*

使用自定义模块
https://go.dev/doc/faq#git_https


安装gin插件  Gin 是一个专注于性能的 Go Web 框架。Gin 的特点是极简主义，适用于构建小型、专注的应用程序。

go get -u github.com/gin-gonic/gin

Echo是另一个高性能的Go Web框架，以其极简的设计和丰富的功能受到开发者的青睐。Echo的目标是提供一流的开发体验和卓越的性能，同时支持中间件、路由组、数据绑定等高级特性。

import "github.com/labstack/echo/v4"

Go Micro是一个专门为微服务设计的框架，提供了一整套构建、部署和管理微服务的工具和库。Go Micro的设计理念是简化微服务的开发和运维，使得开发者可以专注于业务逻辑，而无需关注底层基础设施。

import ("github.com/micro/go-micro/v2")

选择指南
在选择Go微服务框架时，开发者需要综合考虑项目需求、团队技能水平和性能要求。以下是一些选择指南：

Gin：适合需要高性能和简洁API的项目，特别是小型和快速迭代的应用。

Echo：适合需要高性能和丰富功能的项目，特别是中大型和企业级应用。

Go Micro：适合需要完整微服务架构解决方案的项目，特别是复杂的分布式系统和大规模微服务架构。

http://m1.topgoer.com/Components/API.html

Encore是构建Go微服务的现代替代方案，旨在使构建面向事件驱动和分布式系统变得更简单。Encore解决了本地开发体验，并通过使用来自AWS和GCP的强大可扩展服务协助部署。它通过提供微服务框架和基础设施SDK来工作，让您将基础设施语义声明为应用程序代码的一部分。官方文档：https://encore.dev/。

这里是本文中包含的不同框架的高级概述，这些框架根据对各种用例和功能的内置支持水平进行比较。

https://encore.dev/go

|	| Encore |	GoMicro |	Go Kit  |	Gin |
| :------------- | :----------: | :----------: | :----------: | :----------: |
|	描述	|	微服务框架和基础设施 SDK	|	微服务框架	|	Go库集合	|	Web框架 |	
|	使用案例	|	事件驱动、分布式系统	|	事件驱动、分布式系统	|	分布式系统	|	高性能的HTTP API
|	支持协议	|	HTTP	|	HTTP,gRPC	|	HTTP,gRPC	|	HTTP
|	低模板	|	✅︎	|	✅︎	|	✅︎	|	✅︎
|	认证	|	✅︎	|	✅︎	|	✅︎	|	✅︎
|	服务发现	|	✅︎	|	✅︎	|	✅︎	|	❌
|	异步消息	|	✅︎	|	✅︎	|	❌	|	❌
|	内置API文档	|	✅︎	|	✅︎	|	❌	|	❌
|	自动本地开发基础设施	|	✅︎	|	❌	|	❌	|	❌
|	内置跟踪	|	✅︎	|	❌	|	❌	|	❌
|	内置架构图	|	✅︎	|	❌	|	❌	|	❌
|	内置的秘密管理	|	✅︎	|	❌	|	❌	|	❌
|	内置预览环境	|	✅︎	|	❌	|	❌	|	❌
|	内置云基础设施自动化	|	✅︎	|	❌	|	❌	|	❌
|	内置云成本分析	|	✅︎	|	❌	|	❌	|	❌


Go Zero:一个集成了各种工程实践的 web 和 rpc 框架。通过弹性设计保障了大并发服务端的稳定性，经受了充分的实战检验。

https://go-zero.dev/

go zero 安装后，启动demo api 还需要以下插件，整个demo文件夹内的内容最好移动到根目录
```shell
#除了已经提示的那些要安装的插件外，还需要安装这些
go get -u google.golang.org/grpc@latest
go get -u google.golang.org/genproto
go get go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc@v1.24.0
```
#### 修改模板：
https://www.cnblogs.com/wangzhongyang/p/18181397

```shell
#生成api模板，为防止代码被修改，一般情况下，我们会生成到一个临时目录
goctl api go -api website.api -dir ./temp
#生成grpc模板
goctl rpc protoc --go_out=. --go-grpc_out=. --zrpc_out=. ./etcd/server.proto
```
#### 验证器
https://www.jb51.net/jiaoben/31416832z.htm

#### etcd
https://blog.csdn.net/qq_34556414/article/details/113841815
```shell
ETCD_VER=v3.4.34

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}

rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /upload/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /upload/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /upload/etcd --strip-components=1
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

/upload/etcd/etcd --version
/upload/etcd/etcdctl version
```
https://blog.csdn.net/m0_73880436/article/details/142407314
```shell
#项目内需要安装插件
go get go.etcd.io/etcd/clientv3
```
#启动单机
```shell
./etcd --listen-client-urls 'http://0.0.0.0:2379' --advertise-client-urls 'http://0.0.0.0:2379'
```

#### i18n
https://blog.csdn.net/gitblog_00014/article/details/136775274
```shell
go get github.com/nicksnyder/go-i18n/v2@latest
```

```text
admin123 0192023a7bbd73250516f069df18b500
```

#### 问题一：
该版本的 test.exe 与你运行的 Windows 版本不兼容。请查看计算机的系统信息，然后联系软件发布者

原因：入口的包名 需要是main