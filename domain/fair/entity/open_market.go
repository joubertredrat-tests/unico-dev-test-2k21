package entity

type OpenMarket struct {
	RegistryID       string
	Name             string
	Latitude         string
	Longitude        string
	SetCens          string
	District         District
	SubCityHall      SubCityHall
	Address          Address
	AddressReference string
}
