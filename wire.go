//go:build wireinject
// +build wireinject

package main

import (
	"crab-dev/simple-go-gin/configs"
	"crab-dev/simple-go-gin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func AppSet(router *gin.Engine) (*Application, error) {
	return &Application{
		router,
	}, nil
}

var providerSet = wire.NewSet(
	configs.CacheSet,
	configs.RouterSet,
	controllers.ControllerSet,
	AppSet,
)

func CreateApp() (*Application, error) {
	panic(wire.Build(providerSet))
}
