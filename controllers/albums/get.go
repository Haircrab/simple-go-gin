package albums

import (
	"crab-dev/simple-go-gin/domains"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAlbums(rg *gin.RouterGroup) {
	rg.GET("/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var albums = domains.Album{ID: id, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99}

		ctx.JSON(http.StatusOK, albums)
	})
}
