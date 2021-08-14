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
	openMarketCreated, err := s.repository.Create(openMarket)
	if err != nil {
		return nil, OpenMarketServiceHoustonError
	}

	return openMarketCreated, nil
}

func (s OpenMarketService) Update(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
	openMarketUpdated, err := s.repository.Update(openMarket)
	if err != nil {
		return nil, OpenMarketServiceHoustonError
	}

	return openMarketUpdated, nil
}

func (s OpenMarketService) Delete(openMarket entity.OpenMarket) error {
	err := s.repository.Delete(openMarket)
	if err != nil {
		return OpenMarketServiceHoustonError
	}

	return nil
}

func (s OpenMarketService) GetByRegistryID(RegistryID string) (*entity.OpenMarket, error) {
	openMarketFound, err := s.repository.GetByRegistryID(RegistryID)
	if err != nil {
		return nil, OpenMarketServiceHoustonError
	}

	return openMarketFound, nil
}
