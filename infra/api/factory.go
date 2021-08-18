package api

import (
	"strconv"

	domainEntity "github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/entity"
)

func createOpenMarkeSearchCriteriatFromListRequest(request OpenMarketListSearchCriteria) domainEntity.OpenMarketSearchCriteria {
	return domainEntity.OpenMarketSearchCriteria{
		DistrictName:        request.DistrictName,
		SubCityHallRegion5:  request.SubCityHallRegion5,
		OpenMarketName:      request.OpenMarketName,
		AddressNeighborhood: request.AddressNeighborhood,
	}
}

func createOpenMarketFromCreateRequest(request OpenMarketCreateRequest) domainEntity.OpenMarket {
	return domainEntity.OpenMarket{
		RegistryID: request.RegistryID,
		Name:       request.Name,
		Latitude:   request.Latitude,
		Longitude:  request.Longitude,
		SetCens:    request.SetCens,
		AreaP:      request.AreaP,
		Address: domainEntity.Address{
			Street:       request.AddressStreet,
			Number:       request.AddressNumber,
			Neighborhood: request.AddressNeighborhood,
		},
		AddressReference: request.AddressReference,
		District: domainEntity.District{
			Code: strconv.FormatUint(request.DistrictCode, 10),
			Name: request.DistrictName,
		},
		SubCityHall: domainEntity.SubCityHall{
			Code:    strconv.FormatUint(request.SubCityHallCode, 10),
			Name:    request.SubCityHallName,
			Region5: request.SubCityHallRegion5,
			Region8: request.SubCityHallRegion8,
		},
	}
}

func createOpenMarketFromUpdateRequest(request OpenMarketUpdateRequest) domainEntity.OpenMarket {
	return domainEntity.OpenMarket{
		Name:      request.Name,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		SetCens:   request.SetCens,
		AreaP:     request.AreaP,
		Address: domainEntity.Address{
			Street:       request.AddressStreet,
			Number:       request.AddressNumber,
			Neighborhood: request.AddressNeighborhood,
		},
		AddressReference: request.AddressReference,
		District: domainEntity.District{
			Code: strconv.FormatUint(request.DistrictCode, 10),
			Name: request.DistrictName,
		},
		SubCityHall: domainEntity.SubCityHall{
			Code:    strconv.FormatUint(request.SubCityHallCode, 10),
			Name:    request.SubCityHallName,
			Region5: request.SubCityHallRegion5,
			Region8: request.SubCityHallRegion8,
		},
	}
}

func createResponseFromOpenMarket(openMarket domainEntity.OpenMarket) OpenMarketResponse {
	return OpenMarketResponse{
		RegistryID: openMarket.RegistryID,
		Name:       openMarket.Name,
		Latitude:   openMarket.Latitude,
		Longitude:  openMarket.Longitude,
		SetCens:    openMarket.SetCens,
		AreaP:      openMarket.AreaP,
		Address: struct {
			Street       string `json:"street"`
			Number       string `json:"number"`
			Neighborhood string `json:"neighborhood"`
		}{
			Street:       openMarket.Address.Street,
			Number:       openMarket.Address.Number,
			Neighborhood: openMarket.Address.Neighborhood,
		},
		AddressReference: openMarket.AddressReference,
		District: struct {
			Code string `json:"code"`
			Name string `json:"name"`
		}{
			Code: openMarket.District.Code,
			Name: openMarket.District.Name,
		},
		SubCityHall: struct {
			Code    string `json:"code"`
			Name    string `json:"name"`
			Region5 string `json:"region5"`
			Region8 string `json:"region8"`
		}{
			Code:    openMarket.SubCityHall.Code,
			Name:    openMarket.SubCityHall.Name,
			Region5: openMarket.SubCityHall.Region5,
			Region8: openMarket.SubCityHall.Region8,
		},
		CreatedAt: openMarket.CreatedAt,
		UpdatedAt: openMarket.UpdatedAt,
	}
}

func createResponseFromOpenMarketList(openMarketList []*domainEntity.OpenMarket) []OpenMarketResponse {
	var openMarketResponse []OpenMarketResponse

	for _, openMarket := range openMarketList {
		openMarketResponse = append(openMarketResponse, createResponseFromOpenMarket(*openMarket))
	}

	return openMarketResponse
}
