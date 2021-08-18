package worker

import (
	domainEntity "github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/entity"
)

const (
	COLUMN_REGISTRY_ID           = 12
	COLUMN_NAME                  = 11
	COLUMN_LATITUDE              = 2
	COLUMN_LONGITUDE             = 1
	COLUMN_SETCENS               = 3
	COLUMN_AREAP                 = 4
	COLUMN_ADDRESS_STREET        = 13
	COLUMN_ADDRESS_NUMBER        = 14
	COLUMN_ADDRESS_NEIGHBORHOOD  = 15
	COLUMN_ADDRESS_REFERENCE     = 16
	COLUMN_DISTRICT_CODE         = 5
	COLUMN_DISTRICT_NAME         = 6
	COLUMN_SUB_CITY_HALL_CODE    = 7
	COLUMN_SUB_CITY_HALL_NAME    = 8
	COLUMN_SUB_CITY_HALL_REGION5 = 9
	COLUMN_SUB_CITY_HALL_REGION8 = 10
)

func createOpenMarketFromCSVLine(line []string) domainEntity.OpenMarket {
	return domainEntity.OpenMarket{
		RegistryID: line[COLUMN_REGISTRY_ID],
		Name:       line[COLUMN_NAME],
		Latitude:   line[COLUMN_LATITUDE],
		Longitude:  line[COLUMN_LONGITUDE],
		SetCens:    line[COLUMN_SETCENS],
		AreaP:      line[COLUMN_AREAP],
		Address: domainEntity.Address{
			Street:       line[COLUMN_ADDRESS_STREET],
			Number:       line[COLUMN_ADDRESS_NUMBER],
			Neighborhood: line[COLUMN_ADDRESS_NEIGHBORHOOD],
		},
		AddressReference: line[COLUMN_ADDRESS_REFERENCE],
		District: domainEntity.District{
			Code: line[COLUMN_DISTRICT_CODE],
			Name: line[COLUMN_DISTRICT_NAME],
		},
		SubCityHall: domainEntity.SubCityHall{
			Code:    line[COLUMN_SUB_CITY_HALL_CODE],
			Name:    line[COLUMN_SUB_CITY_HALL_NAME],
			Region5: line[COLUMN_SUB_CITY_HALL_REGION5],
			Region8: line[COLUMN_SUB_CITY_HALL_REGION8],
		},
	}
}
