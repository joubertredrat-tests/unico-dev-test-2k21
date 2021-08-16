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
			name: "Test create open market with unknown error on repository",
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
		{
			name: "Test create open market with already exists error on repository",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeCreate = func(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
					return nil, repository.OpenMarketRepositoryAlreadyExistError
				}

				return repo
			}(),
			openMarketExpected: nil,
			errorExpected:      service.OpenMarketServiceAlreadyExistError,
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
			name: "Test update open market with unknown error on repository",
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
		{
			name: "Test update open market with not found error on repository",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeUpdate = func(openMarket entity.OpenMarket) (*entity.OpenMarket, error) {
					return nil, repository.OpenMarketRepositoryNotFoundError
				}

				return repo
			}(),
			openMarketExpected: nil,
			errorExpected:      service.OpenMarketServiceNotFoundError,
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

func TestOpenMarketServiceDelete(t *testing.T) {
	tests := []struct {
		name          string
		repository    repository.OpenMarketRepository
		errorExpected error
	}{
		{
			name: "Test delete open market with success",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeDelete = func(RegistryID string) error {
					return nil
				}

				return repo
			}(),
			errorExpected: nil,
		},
		{
			name: "Test delete open market with open market not found error on repository",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeDelete = func(RegistryID string) error {
					return repository.OpenMarketRepositoryNotFoundError
				}

				return repo
			}(),
			errorExpected: service.OpenMarketServiceNotFoundError,
		},
		{
			name: "Test delete open market with unknown error on repository",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeDelete = func(RegistryID string) error {
					return repository.OpenMarketRepositoryHoustonError
				}

				return repo
			}(),
			errorExpected: service.OpenMarketServiceHoustonError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			openMarketService := service.NewOpenMarketService(test.repository)
			errGot := openMarketService.Delete("4041-0")

			assert.Equal(t, test.errorExpected, errGot)
		})
	}
}

func TestOpenMarketServiceGetByRegistryID(t *testing.T) {
	tests := []struct {
		name               string
		repository         repository.OpenMarketRepository
		openMarketExpected *entity.OpenMarket
		errorExpected      error
	}{
		{
			name: "Test get open market by registry id with success",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeGetByRegistryID = func(RegistryID string) (*entity.OpenMarket, error) {
					openMarketFound := &entity.OpenMarket{
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

					return openMarketFound, nil
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
			name: "Test get open market by registry id with open market not found error on repository",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeGetByRegistryID = func(RegistryID string) (*entity.OpenMarket, error) {
					return nil, repository.OpenMarketRepositoryNotFoundError
				}

				return repo
			}(),
			openMarketExpected: nil,
			errorExpected:      service.OpenMarketServiceNotFoundError,
		},
		{
			name: "Test get open market by registry id with unknown error on repository",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeGetByRegistryID = func(RegistryID string) (*entity.OpenMarket, error) {
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

			openMarketGot, errGot := openMarketService.GetByRegistryID("4041-0")

			assert.Equal(t, test.openMarketExpected, openMarketGot)
			assert.Equal(t, test.errorExpected, errGot)
		})
	}
}

func TestOpenMarketServiceGetListByCriteria(t *testing.T) {
	tests := []struct {
		name                   string
		repository             repository.OpenMarketRepository
		openMarketListExpected []*entity.OpenMarket
		errorExpected          error
	}{
		{
			name: "Test get open market list by criteria with success",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeGetListByCriteria = func(searchCriteria entity.OpenMarketSearchCriteria) ([]*entity.OpenMarket, error) {
					openMarketList := []*entity.OpenMarket{
						&entity.OpenMarket{
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
						&entity.OpenMarket{
							RegistryID: "4045-2",
							Name:       "PRACA SANTA HELENA",
							Latitude:   "-23584852",
							Longitude:  "-46574716",
							SetCens:    "355030893000035",
							AreaP:      "3550308005042",
							Address: entity.Address{
								Street:       "RUA JOSE DOS REIS",
								Number:       "909.000000",
								Neighborhood: "VL ZELINA",
							},
							AddressReference: "RUA OLIVEIRA GOUVEIA",
							District: entity.District{
								Code: "95",
								Name: "VILA PRUDENTE",
							},
							SubCityHall: entity.SubCityHall{
								Code:    "29",
								Name:    "VILA PRUDENTE",
								Region5: "Leste",
								Region8: "Leste 1",
							},
							CreatedAt: getDateMock("2021-08-14T16:37:05-0300"),
							UpdatedAt: getDateMock("2021-08-14T16:41:05-0300"),
						},
					}

					return openMarketList, nil
				}

				return repo
			}(),
			openMarketListExpected: []*entity.OpenMarket{
				&entity.OpenMarket{
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
				&entity.OpenMarket{
					RegistryID: "4045-2",
					Name:       "PRACA SANTA HELENA",
					Latitude:   "-23584852",
					Longitude:  "-46574716",
					SetCens:    "355030893000035",
					AreaP:      "3550308005042",
					Address: entity.Address{
						Street:       "RUA JOSE DOS REIS",
						Number:       "909.000000",
						Neighborhood: "VL ZELINA",
					},
					AddressReference: "RUA OLIVEIRA GOUVEIA",
					District: entity.District{
						Code: "95",
						Name: "VILA PRUDENTE",
					},
					SubCityHall: entity.SubCityHall{
						Code:    "29",
						Name:    "VILA PRUDENTE",
						Region5: "Leste",
						Region8: "Leste 1",
					},
					CreatedAt: getDateMock("2021-08-14T16:37:05-0300"),
					UpdatedAt: getDateMock("2021-08-14T16:41:05-0300"),
				},
			},
			errorExpected: nil,
		},
		{
			name: "Test get open market list by criteria with error on repository",
			repository: func() repository.OpenMarketRepository {
				repo := repository.NewOpenMarketRepositoryFake()
				repo.FakeGetListByCriteria = func(searchCriteria entity.OpenMarketSearchCriteria) ([]*entity.OpenMarket, error) {
					return nil, repository.OpenMarketRepositoryHoustonError
				}

				return repo
			}(),
			openMarketListExpected: nil,
			errorExpected:          service.OpenMarketServiceHoustonError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			openMarketService := service.NewOpenMarketService(test.repository)
			searchCriteria := entity.NewOpenMarketSearchCriteria()
			searchCriteria.SubCityHallRegion5 = "Leste"

			openMarketListGot, errGot := openMarketService.GetListByCriteria(searchCriteria)

			assert.Equal(t, test.openMarketListExpected, openMarketListGot)
			assert.Equal(t, test.errorExpected, errGot)
		})
	}
}

func getDateMock(date string) *time.Time {
	t, _ := time.Parse("2006-01-02T15:04:05-0300", date)

	return &t
}
