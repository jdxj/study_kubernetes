package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.GET("/", hello)
	err := engine.Run(":8080")
	if err != nil {
		log.Printf("Run: %s\n", err)
	}
}

func hello(ctx *gin.Context) {
	name, err := os.Hostname()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"hostname": name})
}
