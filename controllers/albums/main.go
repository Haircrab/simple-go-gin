package albums

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type DetailsController struct {
	// logger  *zap.Logger
	// service services.DetailsService
}

func NewDetailsController() *DetailsController {
	return &DetailsController{
		// logger:  logger,
		// service: s,
	}
}

type InitControllers func(r *gin.Engine)

func CreateInitControllersFn(
	pc *DetailsController,
) InitControllers {
	return func(r *gin.Engine) {
		route := r.Group("/albums")
		{
			GetAlbums(route)
			ListAlbums(route)
		}
	}
}

var ControllerSet = wire.NewSet(NewDetailsController, CreateInitControllersFn)
