package albums

import (
	"crab-dev/simple-go-gin/albums/domains"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (pc *AlbumsController) List(ctx *gin.Context) {
	var albums = []domains.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	ctx.JSON(http.StatusOK, albums)
}
