package main

import (
	"flag"

	"github.com/DeleteElf/zero-user/middleware"
	"github.com/DeleteElf/zero-user/rpchandles/character"
	"github.com/DeleteElf/zero-user/rpchandles/user"
	"github.com/DeleteElf/zero-user/server"

	//"github.com/DeleteElf/zero-user/rest/handler/character"
	//"github.com/DeleteElf/zero-user/rpchandles/user"
	//"github.com/DeleteElf/zero-user/rest/handler/user"
	"github.com/DeleteElf/zero-web-core/service"
	"github.com/deleteelf/zero-core/utils/loghelper"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/config.yaml", "主配置文件")

func main() {
	flag.Parse() //读取配置文件
	var serviceCfg service.ServiceConf
	conf.MustLoad(*configFile, &serviceCfg) //解析配置

	switch serviceCfg.Service {
	case service.RpcClient:
		var cfg service.RpcClientConf
		conf.MustLoad(*configFile, &cfg)
		s := rest.MustNewServer(cfg.RestConf) //初始化服务
		sc := service.NewServiceContext(&cfg)

		sc.Middlewares["Auth"] = middleware.NewAuthInterceptor(sc.RedisClient).Handle
		if cfg.Plugins["user"] {
			//user.RegisterHandlers(s, sc)
		}
		if cfg.Plugins["character"] {
			//character.RegisterHandlers(s, sc) //多个服务要分别注册
		}
		defer s.Stop()
		loghelper.GetDefaultLogger().Infof("rpc 客户端启动！！！监听地址%s:%d...", cfg.Host, cfg.Port)
		s.Start()
	case service.RpcServer:
		var cfg service.RpcServerConf
		conf.MustLoad(*configFile, &cfg)
		loghelper.GetLogManager().Init(loghelper.LogLevel(cfg.DbConfig.LogLevel)) //初始化数据库专属日志系统
		loghelper.GetDefaultLogger().Info("数据库日志系统启动！！！")
		sc := service.NewServiceContext(cfg)
		s := zrpc.MustNewServer(cfg.RpcServerConf, func(grpcServer *grpc.Server) {
			//在这里进行多个服务的实际注册
			if cfg.Plugins["user"] {
				user.RegisterUserServiceServer(grpcServer, server.NewUserServiceServer(sc))
			}
			if cfg.Plugins["character"] {
				character.RegisterCharacterServiceServer(grpcServer, server.NewCharacterServiceServer(sc))
			}
			//if cfg.Mode == service.DevMode || cfg.Mode == service.TestMode {
			//	reflection.Register(grpcServer)
			//}
		})
		defer s.Stop()
		loghelper.GetDefaultLogger().Infof("rpc服务启动！！！监听地址%s...", cfg.ListenOn)
		s.Start()
	case service.Rest:
		var cfg service.RestConf
		conf.MustLoad(*configFile, &cfg)
		loghelper.GetLogManager().Init(loghelper.LogLevel(cfg.DbConfig.LogLevel)) //初始化数据库专属日志系统
		loghelper.GetDefaultLogger().Info("数据库日志系统启动！！！")
		s := rest.MustNewServer(cfg.RestConf) //初始化服务
		sc := service.NewServiceContext(cfg)
		sc.Middlewares["auth"] = middleware.NewAuthInterceptor(sc.RedisClient).Handle
		//if cfg.Plugins["user"] {
		//	user.RegisterHandlers(s, sc)
		//}
		//if cfg.Plugins["character"] {
		//	character.RegisterHandlers(s, sc) //多个服务要分别注册
		//}

		defer s.Stop()
		loghelper.GetDefaultLogger().Infof("网站启动！！！监听地址%s:%d...", cfg.Host, cfg.Port)
		s.Start()
	}
}

//func loadPlugins(service *rest.Server, cfg config.Config) {
//	sc := service.NewServiceContext(cfg)
//	sc.InitPlugin(service, nil) //这里初始化插件，如果配置没有启用，则不会实际加载插件
//}
