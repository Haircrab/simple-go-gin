package albums

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (pc *AlbumsController) List(ctx *gin.Context) {
	var albums, err = pc.service.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "List albums error")
	}

	ctx.JSON(http.StatusOK, albums)
}
