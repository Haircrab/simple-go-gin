package controllers

import (
	"crab-dev/simple-go-gin/albums/controllers/albums"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	albums.ControllerSet,
)
