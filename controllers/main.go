package controllers

import (
	"crab-dev/simple-go-gin/controllers/albums"
	"github.com/google/wire"
)

var ControllerSet = wire.NewSet(
	albums.ControllerSet,
)
