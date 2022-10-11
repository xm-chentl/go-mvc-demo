package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/xm-chentl/go-mvc-demo/api"
	"github.com/xm-chentl/go-mvc-demo/plugin/apisvc"
	"github.com/xm-chentl/go-mvc-demo/plugin/cmuxex"
	"github.com/xm-chentl/go-mvc-demo/proto/go/userservice"
	"github.com/xm-chentl/go-mvc-demo/service/configsvc"
	"github.com/xm-chentl/go-mvc-demo/service/ginsvc"
	"github.com/xm-chentl/go-mvc-demo/service/grpcsvc"
	"google.golang.org/grpc"

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
		if err := apisvc.GenerateMatedata(); err != nil {
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

	ginS := gin.Default()
	ginsvc.NewBgPost(ginS)
	ginsvc.NewMobilePost(ginS)
	ginsvc.NewGraphqlGet(ginS) // todo: 生产环境不开放
	ginsvc.NewGraphqlPost(ginS)
	grpcS := grpc.NewServer()
	userservice.RegisterUserServiceServer(grpcS, grpcsvc.NewUserService())
	s := cmuxex.IntegratedServer{
		GinS:  ginS,
		GrpcS: grpcS,
	}
	if err := s.Run(fmt.Sprintf(":%d", config.Port)); err != nil {
		panic(err)
	}
}
