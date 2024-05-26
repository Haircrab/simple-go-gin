package albums

import (
	"crab-dev/simple-go-gin/albums/services/albums"
	"crab-dev/simple-go-gin/pkg/transports/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type AlbumsController struct {
	logger  *zap.Logger
	service albums.AlbumsService
}

func NewAlbumsController(logger *zap.Logger, s albums.AlbumsService) *AlbumsController {
	return &AlbumsController{
		logger:  logger,
		service: s,
	}
}

func CreateInitControllersFn(pc *AlbumsController) http.InitControllers {
	return func(r *gin.Engine) {
		fmt.Println("initing albums controller")
		r.GET("/albums/:id", pc.Get)
		r.GET("/albums", pc.List)
	}
}

var ControllerSet = wire.NewSet(NewAlbumsController, CreateInitControllersFn)
