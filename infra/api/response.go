package api

import "time"

type OpenMarketResponse struct {
	RegistryID string `json:"registry_id"`
	Name       string `json:"name"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	SetCens    string `json:"set_cens"`
	AreaP      string `json:"area_p"`
	Address    struct {
		Street       string `json:"street"`
		Number       string `json:"number"`
		Neighborhood string `json:"neighborhood"`
	} `json:"address"`
	AddressReference string `json:"address_reference"`
	District         struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"district"`
	SubCityHall struct {
		Code    string `json:"code"`
		Name    string `json:"name"`
		Region5 string `json:"region5"`
		Region8 string `json:"region8"`
	} `json:"sub_city_hall"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
