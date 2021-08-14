package repository

import (
	"github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/entity"
)

type OpenMarketRepositoryFake struct {
	FakeCreate            func(openMarket entity.OpenMarket) (*entity.OpenMarket, error)
	FakeUpdate            func(openMarket entity.OpenMarket) (*entity.OpenMarket, error)
	FakeDelete            func(openMarket entity.OpenMarket) error
	FakeGetByRegistryID   func(RegistryID string) (*entity.OpenMarket, error)
	FakeGetListByCriteria func(searchCriteria entity.OpenMarketSearchCriteria) ([]*entity.OpenMarket, error)
}

func NewOpenMarketRepositoryFake() OpenMarketRepositoryFake {
	return OpenMarketRepositoryFake{}
}

func (r OpenMarketRepositoryFake) Create(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
	return r.FakeCreate(openMarket)
}

func (r OpenMarketRepositoryFake) Update(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
	return r.FakeUpdate(openMarket)
}

func (r OpenMarketRepositoryFake) Delete(openMarket entity.OpenMarket) error {
	return r.FakeDelete(openMarket)
}

func (r OpenMarketRepositoryFake) GetByRegistryID(RegistryID string) (*entity.OpenMarket, error) {
	return r.FakeGetByRegistryID(RegistryID)
}

func (r OpenMarketRepositoryFake) GetListByCriteria(searchCriteria entity.OpenMarketSearchCriteria) ([]*entity.OpenMarket, error) {
	return r.FakeGetListByCriteria(searchCriteria)
}
