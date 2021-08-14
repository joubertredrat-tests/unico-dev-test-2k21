package service_test

import (
	"testing"
	"time"

	"github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/entity"
	"github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/repository"
	"github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/service"
	"github.com/stretchr/testify/assert"
)

func TestOpenMarketServiceCreate(t *testing.T) {
	tests := []struct {
		name               string
		repository         repository.OpenMarketRepository
		openMarketExpected *entity.OpenMarket
		errorExpected      error
	}{
		{
			name: "Test create open market with success",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeCreate = func(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
					openMarketCreated := &entity.OpenMarket{
						RegistryID: "4041-0",
						Name:       "VILA FORMOSA",
						Latitude:   "-23558733",
						Longitude:  "-46550164",
						SetCens:    "355030885000091",
						AreaP:      "3550308005040",
						Address: entity.Address{
							Street:       "RUA MARAGOJIPE",
							Number:       "S/N",
							Neighborhood: "VL FORMOSA",
						},
						AddressReference: "TV RUA PRETORIA",
						District: entity.District{
							Code: "87",
							Name: "VILA FORMOSA",
						},
						SubCityHall: entity.SubCityHall{
							Code:    "26",
							Name:    "ARICANDUVA-FORMOSA-CARRAO",
							Region5: "Leste",
							Region8: "Leste 1",
						},
						CreatedAt: getDateMock("2021-08-14T11:15:05-0300"),
					}

					return openMarketCreated, nil
				}

				return repo
			}(),
			openMarketExpected: &entity.OpenMarket{
				RegistryID: "4041-0",
				Name:       "VILA FORMOSA",
				Latitude:   "-23558733",
				Longitude:  "-46550164",
				SetCens:    "355030885000091",
				AreaP:      "3550308005040",
				Address: entity.Address{
					Street:       "RUA MARAGOJIPE",
					Number:       "S/N",
					Neighborhood: "VL FORMOSA",
				},
				AddressReference: "TV RUA PRETORIA",
				District: entity.District{
					Code: "87",
					Name: "VILA FORMOSA",
				},
				SubCityHall: entity.SubCityHall{
					Code:    "26",
					Name:    "ARICANDUVA-FORMOSA-CARRAO",
					Region5: "Leste",
					Region8: "Leste 1",
				},
				CreatedAt: getDateMock("2021-08-14T11:15:05-0300"),
			},
			errorExpected: nil,
		},
		{
			name: "Test create open market with error on repository",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeCreate = func(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
					return nil, repository.OpenMarketRepositoryHoustonError
				}

				return repo
			}(),
			openMarketExpected: nil,
			errorExpected:      service.OpenMarketServiceHoustonError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			openMarketService := service.NewOpenMarketService(test.repository)

			openMarketGot, errGot := openMarketService.Create(entity.OpenMarket{
				RegistryID: "4041-0",
				Name:       "VILA FORMOSA",
				Latitude:   "-23558733",
				Longitude:  "-46550164",
				SetCens:    "355030885000091",
				AreaP:      "3550308005040",
				Address: entity.Address{
					Street:       "RUA MARAGOJIPE",
					Number:       "S/N",
					Neighborhood: "VL FORMOSA",
				},
				AddressReference: "TV RUA PRETORIA",
				District: entity.District{
					Code: "87",
					Name: "VILA FORMOSA",
				},
				SubCityHall: entity.SubCityHall{
					Code:    "26",
					Name:    "ARICANDUVA-FORMOSA-CARRAO",
					Region5: "Leste",
					Region8: "Leste 1",
				},
			})

			assert.Equal(t, test.openMarketExpected, openMarketGot)
			assert.Equal(t, test.errorExpected, errGot)
		})
	}
}

func TestOpenMarketServiceUpdate(t *testing.T) {
	tests := []struct {
		name               string
		repository         repository.OpenMarketRepository
		openMarketExpected *entity.OpenMarket
		errorExpected      error
	}{
		{
			name: "Test update open market with success",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeUpdate = func(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
					openMarketUpdated := &entity.OpenMarket{
						RegistryID: "4041-0",
						Name:       "VILA FORMOSA",
						Latitude:   "-23558733",
						Longitude:  "-46550164",
						SetCens:    "355030885000091",
						AreaP:      "3550308005040",
						Address: entity.Address{
							Street:       "RUA MARAGOJIPE",
							Number:       "S/N",
							Neighborhood: "VL FORMOSA",
						},
						AddressReference: "TV RUA PRETORIA",
						District: entity.District{
							Code: "87",
							Name: "VILA FORMOSA",
						},
						SubCityHall: entity.SubCityHall{
							Code:    "26",
							Name:    "ARICANDUVA-FORMOSA-CARRAO",
							Region5: "Leste",
							Region8: "Leste 1",
						},
						CreatedAt: getDateMock("2021-08-14T11:15:05-0300"),
						UpdatedAt: getDateMock("2021-08-14T11:21:05-0300"),
					}

					return openMarketUpdated, nil
				}

				return repo
			}(),
			openMarketExpected: &entity.OpenMarket{
				RegistryID: "4041-0",
				Name:       "VILA FORMOSA",
				Latitude:   "-23558733",
				Longitude:  "-46550164",
				SetCens:    "355030885000091",
				AreaP:      "3550308005040",
				Address: entity.Address{
					Street:       "RUA MARAGOJIPE",
					Number:       "S/N",
					Neighborhood: "VL FORMOSA",
				},
				AddressReference: "TV RUA PRETORIA",
				District: entity.District{
					Code: "87",
					Name: "VILA FORMOSA",
				},
				SubCityHall: entity.SubCityHall{
					Code:    "26",
					Name:    "ARICANDUVA-FORMOSA-CARRAO",
					Region5: "Leste",
					Region8: "Leste 1",
				},
				CreatedAt: getDateMock("2021-08-14T11:15:05-0300"),
				UpdatedAt: getDateMock("2021-08-14T11:21:05-0300"),
			},
			errorExpected: nil,
		},
		{
			name: "Test update open market with error on repository",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeUpdate = func(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
					return nil, repository.OpenMarketRepositoryHoustonError
				}

				return repo
			}(),
			openMarketExpected: nil,
			errorExpected:      service.OpenMarketServiceHoustonError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			openMarketService := service.NewOpenMarketService(test.repository)

			openMarketGot, errGot := openMarketService.Update(entity.OpenMarket{
				RegistryID: "4041-0",
				Name:       "VILA FORMOSA",
				Latitude:   "-23558733",
				Longitude:  "-46550164",
				SetCens:    "355030885000091",
				AreaP:      "3550308005040",
				Address: entity.Address{
					Street:       "RUA MARAGOJIPE",
					Number:       "S/N",
					Neighborhood: "VL FORMOSA",
				},
				AddressReference: "TV RUA PRETORIA",
				District: entity.District{
					Code: "87",
					Name: "VILA FORMOSA",
				},
				SubCityHall: entity.SubCityHall{
					Code:    "26",
					Name:    "ARICANDUVA-FORMOSA-CARRAO",
					Region5: "Leste",
					Region8: "Leste 1",
				},
				CreatedAt: getDateMock("2021-08-14T11:15:05-0300"),
			})

			assert.Equal(t, test.openMarketExpected, openMarketGot)
			assert.Equal(t, test.errorExpected, errGot)
		})
	}
}

func getDateMock(date string) *time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05-0300", date)

	return &t
}
