package configs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type InitControllers func(r *gin.Engine)

func InitRouter(init InitControllers) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "OK") })

	init(r)

	return r
}

var RouterSet = wire.NewSet(InitRouter)
