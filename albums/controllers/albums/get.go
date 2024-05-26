package albums

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (pc *AlbumsController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	var albums, err = pc.service.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "Get one albums error")
	}

	ctx.JSON(http.StatusOK, albums)
}
