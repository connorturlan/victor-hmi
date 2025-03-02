package common

// SA_FuelPriceList is the raw json representation of the fuel data.
type SA_FuelPriceList struct {
	Prices []SA_FuelPrice `json:"SitePrices"`
}

// SA_FuelPrice is the raw price data.
type SA_FuelPrice struct {
	SiteId             int     `json:"SiteId"`
	FuelId             int     `json:"FuelId"`
	CollectionMethod   string  `json:"CollectionMethod"`
	TransactionDateUTC string  `json:"TransactionDateUTC"`
	Price              float64 `json:"Price"`
}

// SA_PetrolStationList is the raw json representation of the sites data.
type SA_PetrolStationList struct {
	Sites []SA_PetrolStationSite `json:"S"`
}

// SA_PetrolStationSite is the raw json representation of the sites data.
type SA_PetrolStationSite struct {
	SiteID        int     `json:"S"`
	Address       string  `json:"A"`
	Name          string  `json:"N"`
	BrandID       int     `json:"B"`
	Postcode      string  `json:"P"`
	LastUpdated   string  `json:"D"`
	GooglePlaceID string  `json:"GPI"`
	Latitude      float64 `json:"Lat"`
	Longitude     float64 `json:"Lng"`
}

type FuelPriceList struct {
	Sites map[int]FuelStation `json:"Sites"`
}

type FuelStation struct {
	SiteID      int               `json:"SiteId"`
	LastUpdated string            `json:"LastUpdated"`
	FuelTypes   map[int]FuelPrice `json:"FuelTypes"`
}

type FuelPrice struct {
	FuelID             int    `json:"FuelId"`
	CollectionMethod   string `json:"CollectionMethod"`
	TransactionDateUTC string `json:"TransactionDateUTC"`
	Price              int    `json:"Price"`
}

type FuelStationDetails struct {
	SiteID        int               `json:"SiteId"`
	Name          string            `json:"Name"`
	Lat           float64           `json:"Lat"`
	Lng           float64           `json:"Lng"`
	GooglePlaceID string            `json:"GPI"`
	Address       string            `json:"Address"`
	BrandID       int               `json:"BrandID"`
	Postcode      string            `json:"Postcode"`
	LastUpdated   string            `json:"LastUpdated"`
	FuelTypes     map[int]FuelPrice `json:"FuelTypes"`
}
