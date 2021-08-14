package entity

import "time"

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
	CreatedAt        time.Time
	UpdatedAt        *time.Time
}
