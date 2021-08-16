package service

import (
	"errors"

	"github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/entity"
	"github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/repository"
)

var (
	OpenMarketServiceHoustonError      = errors.New("Anything wrong is not right")
	OpenMarketServiceNotFoundError     = errors.New("Open market informed by registry ID does not exist")
	OpenMarketServiceAlreadyExistError = errors.New("Open market with informed registry ID can not be created, already exists")
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
		if errors.Is(err, repository.OpenMarketRepositoryAlreadyExistError) {
			return nil, OpenMarketServiceAlreadyExistError
		}

		return nil, OpenMarketServiceHoustonError
	}

	return openMarketCreated, nil
}

func (s OpenMarketService) Update(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
	openMarketUpdated, err := s.repository.Update(openMarket)
	if err != nil {
		if errors.Is(err, repository.OpenMarketRepositoryNotFoundError) {
			return nil, OpenMarketServiceNotFoundError
		}

		return nil, OpenMarketServiceHoustonError
	}

	return openMarketUpdated, nil
}

func (s OpenMarketService) Delete(RegistryID string) error {
	err := s.repository.Delete(RegistryID)
	if err != nil {
		if errors.Is(err, repository.OpenMarketRepositoryNotFoundError) {
			return OpenMarketServiceNotFoundError
		}

		return OpenMarketServiceHoustonError
	}

	return nil
}

func (s OpenMarketService) GetByRegistryID(RegistryID string) (*entity.OpenMarket, error) {
	openMarketFound, err := s.repository.GetByRegistryID(RegistryID)
	if err != nil {
		if errors.Is(err, repository.OpenMarketRepositoryNotFoundError) {
			return nil, OpenMarketServiceNotFoundError
		}

		return nil, OpenMarketServiceHoustonError
	}

	return openMarketFound, nil
}

func (s OpenMarketService) GetListByCriteria(searchCriteria entity.OpenMarketSearchCriteria) ([]*entity.OpenMarket, error) {
	openMarketList, err := s.repository.GetListByCriteria(searchCriteria)
	if err != nil {
		return nil, OpenMarketServiceHoustonError
	}

	return openMarketList, nil
}
