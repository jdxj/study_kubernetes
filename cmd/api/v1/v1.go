package v1

import (
	"context"
	"net/http"
	"os"

	echo "github.com/jdxj/study_kubernetes/cmd/echo/proto"

	"github.com/gin-gonic/gin"
)

var (
	EchoService echo.EchoService
)

func Register(p *gin.RouterGroup) {
	gV1 := p.Group("/v1")
	gV1.GET("", hello)

	gV1.POST("/echo", echoHandler)
}

type Message struct {
	Data string `json:"data"`
}

func hello(ctx *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":      "v1",
		"hostname": hostname,
	})
}

func echoHandler(ctx *gin.Context) {
	msg := &Message{}
	err := ctx.BindJSON(msg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}

	req := &echo.Ping{
		Data: []byte(msg.Data),
	}
	resp, err := EchoService.Hello(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": string(resp.Data),
	})
}
