package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/jdxj/study_kubernetes/cmd/api/v1"

	"github.com/gin-gonic/gin"
)

func StartAPI(host, port string) {
	e := gin.Default()

	api := e.Group("/api")
	api.GET("/", hello)

	v1.Register(api)

	err := e.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Printf("Run: %s\n", err)
	}
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "api"})
}
