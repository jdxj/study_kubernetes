package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(p *gin.RouterGroup) {
	gV1 := p.Group("/v1")
	gV1.GET("", hello)
}

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "v1"})
}
