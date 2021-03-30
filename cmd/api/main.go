package main

import (
	"log"

	"github.com/asim/go-micro/v3"
)

const (
	modelName = "api"
)

func main() {
	service := micro.NewService(
		micro.Name(modelName),
	)
	service.Init(
		micro.AfterStart(func() error {
			// todo: 配置
			go StartAPI("127.0.0.1", "8080")
			return nil
		}),
	)
	err := service.Run()
	if err != nil {
		log.Printf("Run: %s\n", err)
	}
}
