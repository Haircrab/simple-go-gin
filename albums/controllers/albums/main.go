package albums

import (
	"crab-dev/simple-go-gin/pkg/transports/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

type AlbumsController struct {
	logger *zap.Logger
	// service services.DetailsService
}

func NewAlbumsController(logger *zap.Logger) *AlbumsController {
	return &AlbumsController{
		logger: logger,
		// service: s,
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
