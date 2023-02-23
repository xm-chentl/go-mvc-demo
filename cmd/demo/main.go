package main

import (
	"flag"
	"os"
	"path"

	"github.com/xm-chentl/go-mvc-demo/api"
	"github.com/xm-chentl/go-mvc-demo/plugin/apisvc"
	"github.com/xm-chentl/go-mvc-demo/service/configsvc"
	"github.com/xm-chentl/go-mvc-demo/service/ginsvc"
	"github.com/xm-chentl/go-mvc-demo/service/grpcsvc"
	"github.com/xm-chentl/gocore/frame"

	"github.com/xm-chentl/go-code/guidex"
	"github.com/xm-chentl/go-code/guidex/snowflake"
	"github.com/xm-chentl/go-mvc/ioc"
)

type Default struct {
	Port  int
	Name  string
	MySQL string
	Mongo string
}

var g = flag.Bool("api", false, "generate api metadata.go")

func main() {
	flag.Parse()
	if *g {
		rootDir, _ := os.Getwd()
		apiDir := path.Join(rootDir, "..", "..", "api")
		if err := apisvc.GenerateMatedata(apiDir); err != nil {
			panic(err)
		}
		return
	}

	// curSrc, _ := os.Getwd()

	var config Default
	cfgSvc := configsvc.NewYaml(path.Join("..", "..", "build", "config.yml"))
	if err := cfgSvc.GetStruct(&config); err != nil {
		panic(err)
	}

	// ioc.Set(new(dbfactory.IDbFactory), dbfactory.New(map[dbtype.Value]dbfactory.IDbFactory{
	// 	dbtype.MySql: gormex.New(config.MySQL),
	// 	dbtype.Mongo: mongoex.New("leban-server", config.Mongo),
	// }))
	ioc.Set(new(guidex.IGenerate), snowflake.New())
	api.Register()

	// ginS := gin.Default()
	// ginsvc.NewBgPost(ginS)
	// ginsvc.NewMobilePost(ginS)
	// ginsvc.NewGraphqlGet(ginS) // todo: 生产环境不开放
	// ginsvc.NewGraphqlPost(ginS)
	// grpcS := grpc.NewServer()
	// userservice.RegisterUserServiceServer(grpcS, grpcsvc.NewUserService())
	// s := cmuxex.IntegratedServer{
	// 	GinSrv:  ginS,
	// 	GrpcSrv: grpcS,
	// }
	// if err := s.Run(fmt.Sprintf(":%d", config.Port)); err != nil {
	// 	panic(err)
	// }

	s := frame.New()
	s.RegisterGin(
		ginsvc.UseOrigin(),
		ginsvc.NewBgPost(),
		ginsvc.NewGraphqlGet(),
		ginsvc.NewGraphqlPost(),
		ginsvc.NewMobilePost(),
		ginsvc.NewUserCenterPost(),
		ginsvc.NewUPload(),
	)
	s.RegisterGRPC(grpcsvc.NewUserService())
	s.Run(config.Port)
}
