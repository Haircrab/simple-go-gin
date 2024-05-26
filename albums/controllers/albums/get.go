package albums

import (
	"crab-dev/simple-go-gin/albums/domains"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (pc *AlbumsController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	var albums = domains.Album{ID: id, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99}

	ctx.JSON(http.StatusOK, albums)
}
