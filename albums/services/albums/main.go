package albums

import (
	"crab-dev/simple-go-gin/albums/domains"

	"github.com/google/wire"
	"go.uber.org/zap"
)

type AlbumsService interface {
	List() (*[]domains.Album, error)
	Get(id string) (*domains.Album, error)
}

type DefaultAlbumsService struct {
	logger *zap.Logger
	// Repository repositories.DetailsRepository
}

func New(logger *zap.Logger) AlbumsService {
	return &DefaultAlbumsService{
		logger: logger.With(zap.String("type", "DefaultDetailsService")),
		// Repository: Repository,
	}
}

func (s *DefaultAlbumsService) Get(id string) (p *domains.Album, err error) {
	var album = domains.Album{ID: id, Title: "Blue Train jsjs", Artist: "John Coltrane", Price: 56.99}
	return &album, nil
}

func (s *DefaultAlbumsService) List() (p *[]domains.Album, err error) {
	var albums = []domains.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	return &albums, nil
}

var ProviderSet = wire.NewSet(New)
