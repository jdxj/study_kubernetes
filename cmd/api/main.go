package main

import (
	"log"

	"github.com/jdxj/study_kubernetes/config"
	"github.com/jdxj/study_kubernetes/dao/db"
	"github.com/jdxj/study_kubernetes/dao/redis"

	v1 "github.com/jdxj/study_kubernetes/cmd/api/v1"
	echo "github.com/jdxj/study_kubernetes/cmd/echo/proto"

	_ "github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/micro/cli/v2"
)

const (
	modelName = "api"
)

func main() {
	service := micro.NewService(
		micro.Name(modelName),
		micro.Flags(&cli.StringFlag{
			Name:  "cfg",
			Value: "../../config/config.yaml",
		}),
	)

	service.Init(
		micro.Action(func(ctx *cli.Context) error {
			path := ctx.String("cfg")
			return Init(path)

		}),
		micro.AfterStart(func() error {
			apiCfg := config.Ins.API
			go StartAPI(apiCfg.Host, apiCfg.Port)
			return nil
		}),
		micro.AfterStop(func() error {
			// todo: 实现 StopAPI()
			return nil
		}),
	)

	// note: 要在 StartAPI() 启动之前初始化
	v1.EchoService = echo.NewEchoService("echo", service.Client())

	err := service.Run()
	if err != nil {
		log.Printf("Run: %s\n", err)
	}
}

func Init(path string) error {
	err := config.Init(path)
	if err != nil {
		return err
	}

	dbCfg := config.Ins.DB
	err = db.Init(dbCfg.User, dbCfg.Pass, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	if err != nil {
		return err
	}

	redisCfg := config.Ins.Redis
	err = redis.InitClient(redisCfg.Host, redisCfg.Port, redisCfg.Pass, redisCfg.DBNum)
	if err != nil {
		return err
	}
	return nil
}
