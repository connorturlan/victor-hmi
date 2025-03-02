package defaults

import "os"

const (
	Region              string = "ap-southeast-2"
	PricesTableName     string = "current_fuel_prices"
	SitesTableName      string = "safpis_fuel_sites"
	HistoryTableName    string = "historic_fuel_prices"
	UsersTableName      string = "fuel_tool_users"
	UserTokensTableName string = "fuel_tool_tokens"
	BatchSize           int    = 25
	FuelURL             string = "https://fppdirectapi-prod.safuelpricinginformation.com.au"
)

var (
	IsLocal         bool   = os.Getenv("local") == "true"
	IsUpdatingSites bool   = os.Getenv("update_sites") == "true"
	Apikey          string = os.Getenv("api_key")
	Origin          string = os.Getenv("cors_origin_url")
)
