package service

import (
	"errors"

	"github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/entity"
	"github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/repository"
)

var (
	OpenMarketServiceHoustonError = errors.New("Anything wrong is not right")
)

type OpenMarketService struct {
	repository repository.OpenMarketRepository
}

func NewOpenMarketService(repository repository.OpenMarketRepository) OpenMarketService {
	return OpenMarketService{
		repository: repository,
	}
}

func (s OpenMarketService) Create(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
	return s.repository.Create(openMarket)
}

func (s OpenMarketService) Update(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
	return s.repository.Update(openMarket)
}

func (s OpenMarketService) Delete(openMarket entity.OpenMarket) error {
	return s.repository.Delete(openMarket)
}

func (s OpenMarketService) GetByRegistryID(RegistryID string) (*entity.OpenMarket, error) {
	return s.repository.GetByRegistryID(RegistryID)
}
