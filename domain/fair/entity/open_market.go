package entity

import "time"

type OpenMarket struct {
	RegistryID       string
	Name             string
	Latitude         string
	Longitude        string
	SetCens          string
	AreaP            string
	Address          Address
	AddressReference string
	District         District
	SubCityHall      SubCityHall
	CreatedAt        time.Time
	UpdatedAt        *time.Time
}
