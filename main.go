package main

import (
	"flag"

	"github.com/xm-chentl/go-mvc-demo/api"
	"github.com/xm-chentl/go-mvc-demo/plugin/apisvc"
	"github.com/xm-chentl/go-mvc-demo/service/configsvc"
	"github.com/xm-chentl/go-mvc-demo/service/ginsvc"

	"github.com/xm-chentl/go-code/dbfactory"
	"github.com/xm-chentl/go-code/dbfactory/dbtype"
	"github.com/xm-chentl/go-code/gormex"
	"github.com/xm-chentl/go-code/guidex"
	"github.com/xm-chentl/go-code/guidex/snowflake"
	"github.com/xm-chentl/go-code/mongoex"
	"github.com/xm-chentl/go-mvc/ginex"
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

	var config Default
	cfgSvc := configsvc.NewYaml("config.yml")
	if err := cfgSvc.GetStruct(&config); err != nil {
		panic(err)
	}

	ioc.Set(new(dbfactory.IDbFactory), dbfactory.New(map[dbtype.Value]dbfactory.IDbFactory{
		dbtype.MySql: gormex.New(config.MySQL),
		dbtype.Mongo: mongoex.New("leban-server", config.Mongo),
	}))
	ioc.Set(new(guidex.IGenerate), snowflake.New())

	api.Register()
	ginex.New(
		ginsvc.NewBgPost(),
		ginsvc.NewMobilePost(),
		ginsvc.NewGraphqlGet(), // todo: 生产环境不开放
		ginsvc.NewGraphqlPost(),
	).Run(config.Port)
}
