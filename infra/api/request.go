package api

type OpenMarketListSearchCriteria struct {
	DistrictName        string `form:"district_name"`
	SubCityHallRegion5  string `form:"sub_city_hall_region5"`
	OpenMarketName      string `form:"name"`
	AddressNeighborhood string `form:"address_neighborhood"`
}

type OpenMarketCreateRequest struct {
	RegistryID          string `json:"registry_id" validate:"required"`
	Name                string `json:"name" validate:"required"`
	Latitude            string `json:"latitude" validate:"required"`
	Longitude           string `json:"longitude" validate:"required"`
	SetCens             string `json:"set_cens" validate:"required"`
	AreaP               string `json:"area_p" validate:"required"`
	AddressStreet       string `json:"address_street" validate:"required"`
	AddressNumber       string `json:"address_number" validate:"required"`
	AddressNeighborhood string `json:"address_neighborhood" validate:"required"`
	AddressReference    string `json:"address_reference" validate:"required"`
	DistrictCode        uint64 `json:"district_code" validate:"required"`
	DistrictName        string `json:"district_name" validate:"required"`
	SubCityHallCode     uint64 `json:"sub_city_hall_code" validate:"required"`
	SubCityHallName     string `json:"sub_city_hall_name" validate:"required"`
	SubCityHallRegion5  string `json:"sub_city_hall_region5" validate:"required"`
	SubCityHallRegion8  string `json:"sub_city_hall_region8" validate:"required"`
}

type OpenMarketUpdateRequest struct {
	Name                string `json:"name"`
	Latitude            string `json:"latitude"`
	Longitude           string `json:"longitude"`
	SetCens             string `json:"set_cens"`
	AreaP               string `json:"area_p"`
	AddressStreet       string `json:"address_street"`
	AddressNumber       string `json:"address_number"`
	AddressNeighborhood string `json:"address_neighborhood"`
	AddressReference    string `json:"address_reference"`
	DistrictCode        uint64 `json:"district_code"`
	DistrictName        string `json:"district_name"`
	SubCityHallCode     uint64 `json:"sub_city_hall_code"`
	SubCityHallName     string `json:"sub_city_hall_name"`
	SubCityHallRegion5  string `json:"sub_city_hall_region5"`
	SubCityHallRegion8  string `json:"sub_city_hall_region8"`
}
