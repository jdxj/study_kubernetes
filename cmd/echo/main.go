package main

import (
	"log"

	"github.com/jdxj/study_kubernetes/cmd/echo/handler"
	"github.com/jdxj/study_kubernetes/cmd/echo/proto"
	"github.com/jdxj/study_kubernetes/config"
	"github.com/jdxj/study_kubernetes/dao/db"
	"github.com/jdxj/study_kubernetes/dao/redis"

	_ "github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/micro/cli/v2"
)

const (
	modelName = "echo"
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
	)

	err := proto.RegisterEchoHandler(service.Server(), &handler.EchoService{})
	if err != nil {
		log.Fatalf("RegisterEchoHandler: %s\n", err)
	}

	err = service.Run()
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
