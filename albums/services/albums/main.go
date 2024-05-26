package albums

import (
	"crab-dev/simple-go-gin/albums/domains"

	"github.com/google/wire"
)

type AlbumsService interface {
	Get(ID uint64) (*domains.Album, error)
}

type DefaultAlbumsService struct {
	// logger     *zap.Logger
	// Repository repositories.DetailsRepository
}

func New() AlbumsService {
	return &DefaultAlbumsService{
		// logger:  logger.With(zap.String("type","DefaultDetailsService")),
		// Repository: Repository,
	}
}

func (s *DefaultAlbumsService) Get(ID uint64) (p *domains.Album, err error) {

	var album = &domains.Album{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99}

	return album, nil

}

var ProviderSet = wire.NewSet(New)
