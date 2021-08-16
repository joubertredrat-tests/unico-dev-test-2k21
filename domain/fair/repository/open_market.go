package repository

import (
	"errors"

	"github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/entity"
)

var (
	OpenMarketRepositoryHoustonError      = errors.New("Anything wrong is not right")
	OpenMarketRepositoryNotFoundError     = errors.New("Open market not found in open market repository")
	OpenMarketRepositoryAlreadyExistError = errors.New("Open market already exists with registry id in open market repository")
)

type OpenMarketRepository interface {
	Create(openMarket entity.OpenMarket) (*entity.OpenMarket, error)
	Update(openMarket entity.OpenMarket) (*entity.OpenMarket, error)
	Delete(RegistryID string) error
	GetByRegistryID(RegistryID string) (*entity.OpenMarket, error)
	GetListByCriteria(searchCriteria entity.OpenMarketSearchCriteria) ([]*entity.OpenMarket, error)
}
