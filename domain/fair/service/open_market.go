package service

import (
	"errors"
	"fmt"

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
	openMarketGot, err := s.GetByRegistryID(openMarket.RegistryID)
	if err != nil {
		return nil, err
	}

	getValueUpdate := func(valueEntity string, valueRequest interface{}) string {
		str := fmt.Sprintf("%v", valueRequest)
		if str != "" && str != "0" {
			return str
		}

		return valueEntity
	}

	openMarketGot.Name = getValueUpdate(openMarketGot.Name, openMarket.Name)
	openMarketGot.Latitude = getValueUpdate(openMarketGot.Latitude, openMarket.Latitude)
	openMarketGot.Longitude = getValueUpdate(openMarketGot.Longitude, openMarket.Longitude)
	openMarketGot.SetCens = getValueUpdate(openMarketGot.SetCens, openMarket.SetCens)
	openMarketGot.AreaP = getValueUpdate(openMarketGot.AreaP, openMarket.AreaP)
	openMarketGot.Address.Street = getValueUpdate(openMarketGot.Address.Street, openMarket.Address.Street)
	openMarketGot.Address.Number = getValueUpdate(openMarketGot.Address.Number, openMarket.Address.Number)
	openMarketGot.Address.Neighborhood = getValueUpdate(openMarketGot.Address.Neighborhood, openMarket.Address.Neighborhood)
	openMarketGot.AddressReference = getValueUpdate(openMarketGot.AddressReference, openMarket.AddressReference)
	openMarketGot.District.Code = getValueUpdate(openMarketGot.District.Code, openMarket.District.Code)
	openMarketGot.District.Name = getValueUpdate(openMarketGot.District.Name, openMarket.District.Name)
	openMarketGot.SubCityHall.Code = getValueUpdate(openMarketGot.SubCityHall.Code, openMarket.SubCityHall.Code)
	openMarketGot.SubCityHall.Name = getValueUpdate(openMarketGot.SubCityHall.Name, openMarket.SubCityHall.Name)
	openMarketGot.SubCityHall.Region5 = getValueUpdate(openMarketGot.SubCityHall.Region5, openMarket.SubCityHall.Region5)
	openMarketGot.SubCityHall.Region8 = getValueUpdate(openMarketGot.SubCityHall.Region8, openMarket.SubCityHall.Region8)

	openMarketUpdated, err := s.repository.Update(*openMarketGot)
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
