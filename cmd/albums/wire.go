//go:build wireinject
// +build wireinject

package main

import (
	"crab-dev/simple-go-gin/albums"
	"crab-dev/simple-go-gin/albums/controllers"
	"crab-dev/simple-go-gin/albums/services"
	"crab-dev/simple-go-gin/pkg/app"
	"crab-dev/simple-go-gin/pkg/cache"
	"crab-dev/simple-go-gin/pkg/config"
	// "crab-dev/simple-go-gin/pkg/jaeger"
	"crab-dev/simple-go-gin/pkg/logger"
	"crab-dev/simple-go-gin/pkg/transports/http"

	"github.com/google/wire"
)

var servicesSet = wire.NewSet(
	config.ProviderSet,
	logger.ProviderSet,
	// jaeger.ProviderSet,
	cache.ProviderSet,
	http.ProviderSet,
)

var businessServicesSet = wire.NewSet(
	services.ProviderSet,
	controllers.ProviderSet,
)

var appSet = wire.NewSet(servicesSet, businessServicesSet, albums.ProviderSet)

func CreateApp(cf string) (*app.Application, error) {
	panic(wire.Build(appSet))
}
