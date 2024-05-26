package services

import (
	"crab-dev/simple-go-gin/albums/services/albums"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(albums.ProviderSet)
