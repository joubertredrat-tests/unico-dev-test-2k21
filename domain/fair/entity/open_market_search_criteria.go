package entity

type OpenMarketSearchCriteria struct {
	DistrictName        string
	SubCityHallRegion5  string
	OpenMarketName      string
	AddressNeighborhood string
}

func NewOpenMarketSearchCriteria() OpenMarketSearchCriteria {
	return OpenMarketSearchCriteria{}
}
