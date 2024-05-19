package configs

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	return r
}

var RouterSet = wire.NewSet(InitRouter)
