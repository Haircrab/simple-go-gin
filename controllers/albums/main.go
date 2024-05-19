package albums

import (
	"crab-dev/simple-go-gin/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type DetailsController struct {
	// logger  *zap.Logger
	// service services.DetailsService
}

func NewDetailsController() *DetailsController {

	fmt.Println("NewDetailsController")

	return &DetailsController{
		// logger:  logger,
		// service: s,
	}
}

func CreateInitControllersFn(
	pc *DetailsController,
) configs.InitControllers {
	return func(r *gin.Engine) {

		fmt.Println("CreateInitControllersFn")

		r.GET("/albums/:id", pc.Get)
		r.GET("/albums", pc.List)
	}
}

var ControllerSet = wire.NewSet(NewDetailsController, CreateInitControllersFn)
