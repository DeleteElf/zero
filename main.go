package main

import (
	"flag"

	"github.com/DeleteElf/zero-file/client/rpc/fileclient"
	"github.com/DeleteElf/zero-file/rest/routes/fileroute"
	"github.com/DeleteElf/zero-file/rpchandles/file"
	"github.com/DeleteElf/zero-file/server/rpc/fileserver"
	"github.com/DeleteElf/zero-user/client/rpc/characterclient"
	"github.com/DeleteElf/zero-user/client/rpc/userclient"
	"github.com/DeleteElf/zero-user/middleware"
	"github.com/DeleteElf/zero-user/rest/routes/characterroute"
	"github.com/DeleteElf/zero-user/rest/routes/userroute"
	"github.com/DeleteElf/zero-user/rpchandles/character"
	"github.com/DeleteElf/zero-user/rpchandles/user"
	"github.com/DeleteElf/zero-user/server/rpc/charactersever"
	"github.com/DeleteElf/zero-user/server/rpc/userserver"
	"github.com/DeleteElf/zero-web-core/ctx"
	"github.com/zeromicro/go-zero/core/service"
	"google.golang.org/grpc/reflection"

	"github.com/deleteelf/zero-core/utils/loghelper"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/config.yaml", "主配置文件")

func main() {
	flag.Parse() //读取配置文件
	var serviceCfg ctx.ServiceConf
	conf.MustLoad(*configFile, &serviceCfg) //解析配置

	switch serviceCfg.Service {
	case ctx.RpcClient:
		var cfg ctx.RpcClientConf
		conf.MustLoad(*configFile, &cfg)
		s := rest.MustNewServer(cfg.RestConf) //初始化服务
		loghelper.Info("日志系统启动！！！")
		sc := ctx.NewServiceContext(&cfg)
		sc.Middlewares["Auth"] = middleware.NewAuthInterceptor(sc.RedisClient).Handle
		if cfg.Plugins["user"] {
			sc.Services["user"] = ctx.ServiceObject{
				Name:               "UserServicePlugin",
				Version:            "1.0",
				Enabled:            true,
				SupportServiceType: ctx.RpcClient,
				Service:            userclient.NewUserService(zrpc.MustNewClient(cfg.RpcClientConf)), //需要对每个rpc服务进行注册
			}
			userroute.RegisterUserHandlers(s, sc)
		}
		if cfg.Plugins["character"] {
			sc.Services["character"] = ctx.ServiceObject{
				Name:               "CharacterServicePlugin",
				Version:            "1.0",
				Enabled:            true,
				SupportServiceType: ctx.RpcClient,
				Service:            characterclient.NewCharacterService(zrpc.MustNewClient(cfg.RpcClientConf)), //需要对每个rpc服务进行注册
			}
			characterroute.RegisterCharacterHandlers(s, sc)
		}
		if cfg.Plugins["file"] {
			sc.Services["file"] = ctx.ServiceObject{
				Name:               "FileServicePlugin",
				Version:            "1.0",
				Enabled:            true,
				SupportServiceType: ctx.RpcClient,
				Service:            fileclient.NewFileService(zrpc.MustNewClient(cfg.RpcClientConf)), //需要对每个rpc服务进行注册
			}
			fileroute.RegisterFileHandlers(s, sc)
		}
		defer s.Stop()
		loghelper.Infof("rpc 客户端启动！！！监听地址%s:%d...", cfg.Host, cfg.Port)
		s.Start()
	case ctx.RpcServer:
		var cfg ctx.RpcServerConf
		conf.MustLoad(*configFile, &cfg)
		//loghelper.GetLogManager().Init(loghelper.LogLevel(cfg.DbConfig.LogLevel)) //初始化数据库专属日志系统
		sc := ctx.NewServiceContext(cfg)
		loghelper.Info("日志系统启动！！！")
		s := zrpc.MustNewServer(cfg.RpcServerConf, func(grpcServer *grpc.Server) {
			//在这里进行多个服务的实际注册
			if cfg.Plugins["user"] {
				user.RegisterUserServiceServer(grpcServer, userserver.NewUserService(sc))
			}
			if cfg.Plugins["character"] {
				character.RegisterCharacterServiceServer(grpcServer, charactersever.NewCharacterService(sc))
			}
			if cfg.Plugins["file"] {
				file.RegisterFileServiceServer(grpcServer, fileserver.NewFileService(sc))
			}
			if cfg.Mode == service.DevMode || cfg.Mode == service.TestMode {
				reflection.Register(grpcServer)
			}
		})
		defer s.Stop()
		loghelper.Infof("rpc服务启动！！！监听地址%s...", cfg.ListenOn)
		s.Start()
	case ctx.Rest:
		var cfg ctx.RestConf
		conf.MustLoad(*configFile, &cfg)
		//loghelper.GetLogManager().Init(loghelper.LogLevel(cfg.DbConfig.LogLevel)) //初始化数据库专属日志系统

		s := rest.MustNewServer(cfg.RestConf) //初始化服务
		loghelper.Info("日志系统启动！！！")
		sc := ctx.NewServiceContext(cfg)
		sc.Middlewares["auth"] = middleware.NewAuthInterceptor(sc.RedisClient).Handle
		if cfg.Plugins["user"] {
			userroute.RegisterUserHandlers(s, sc)
		}
		if cfg.Plugins["character"] {
			characterroute.RegisterCharacterHandlers(s, sc)
		}
		if cfg.Plugins["file"] {
			fileroute.RegisterFileHandlers(s, sc)
		}
		defer s.Stop()
		loghelper.Infof("网站启动！！！监听地址%s:%d...", cfg.Host, cfg.Port)
		s.Start()
	}
}

//func loadPlugins(service *rest.Server, cfg config.Config) {
//	sc := service.NewServiceContext(cfg)
//	sc.InitPlugin(service, nil) //这里初始化插件，如果配置没有启用，则不会实际加载插件
//}
