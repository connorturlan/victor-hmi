package common

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"fuelpriceservice/pkg/common/defaults"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/shopspring/decimal"
)

func (pricesList SA_FuelPriceList) ToPriceList() (FuelPriceList, error) {
	prices := FuelPriceList{
		Sites: map[int]FuelStation{},
	}

	for _, price := range pricesList.Prices {
		// check the petrol station exists.
		siteId := price.SiteId
		site, ok := prices.Sites[siteId]
		if !ok {
			site = FuelStation{
				SiteID:    siteId,
				FuelTypes: map[int]FuelPrice{},
			}
			prices.Sites[siteId] = site
		}

		// check the fuel record exists.
		fuelId := price.FuelId
		price := FuelPrice{
			FuelID:             fuelId,
			CollectionMethod:   price.CollectionMethod,
			TransactionDateUTC: price.TransactionDateUTC,
			Price:              int(price.Price),
		}
		site.FuelTypes[fuelId] = price
	}

	return prices, nil
}

// FuelPrice.Marshal returns a dynamodb representation of the FuelPrice struct.
func (price FuelPrice) Marshal() (map[string]*dynamodb.AttributeValue, error) {
	return map[string]*dynamodb.AttributeValue{
		"FuelId": {
			N: aws.String(fmt.Sprintf("%d", price.FuelID)),
		},
		"M": {
			S: aws.String(price.CollectionMethod),
		},
		"D": {
			S: aws.String(price.TransactionDateUTC),
		},
		"P": {
			N: aws.String(fmt.Sprintf("%d", price.Price)),
		},
	}, nil
}

func (p *FuelPrice) Unmarshal(record map[string]*dynamodb.AttributeValue) (err error) {
	fuelIdRecord, ok := record["FuelId"]
	if !ok {
		return nil
	}
	p.FuelID, err = strconv.Atoi(*fuelIdRecord.N)
	if err != nil {
		return err
	}

	methodRecord, ok := record["M"]
	if !ok {
		return nil
	}
	p.CollectionMethod = *methodRecord.S

	dateRecord, ok := record["D"]
	if !ok {
		return nil
	}
	p.TransactionDateUTC = *dateRecord.S

	priceRecord, ok := record["P"]
	if !ok {
		return nil
	}
	p.Price, err = strconv.Atoi(*priceRecord.N)
	if err != nil {
		return err
	}

	return nil
}

// FuelPrice.Marshal returns a dynamodb representation of the FuelPrice struct.
func (petrolStation FuelStationDetails) Marshal() (map[string]*dynamodb.AttributeValue, error) {
	return map[string]*dynamodb.AttributeValue{
		"SiteId": {N: aws.String(fmt.Sprintf("%d", petrolStation.SiteID))},
		"A":      {S: aws.String(petrolStation.Address)},
		"N":      {S: aws.String(petrolStation.Name)},
		"B":      {N: aws.String(fmt.Sprintf("%d", petrolStation.BrandID))},
		"P":      {S: aws.String(petrolStation.Postcode)},
		"D":      {S: aws.String(petrolStation.LastUpdated)},
		"G":      {S: aws.String(petrolStation.GooglePlaceID)},
		"Lt":     {N: aws.String(decimal.NewFromFloat(petrolStation.Lat).String())},
		"Lg":     {N: aws.String(decimal.NewFromFloat(petrolStation.Lng).String())},
	}, nil
}

func (petrolStation *FuelStationDetails) Unmarshal(record map[string]*dynamodb.AttributeValue) (err error) {
	petrolStation.Name = SafeS(record, "N")
	petrolStation.GooglePlaceID = SafeS(record, "G")

	SiteId, err := strconv.Atoi(*record["SiteId"].N)
	if err != nil {
		return err
	}
	petrolStation.SiteID = SiteId

	lat, err := strconv.ParseFloat(*record["Lt"].N, 64)
	if err != nil {
		return err
	}
	petrolStation.Lat = lat

	lng, err := strconv.ParseFloat(*record["Lg"].N, 64)
	if err != nil {
		return err
	}
	petrolStation.Lng = lng

	return nil
}

// FuelStation.Marshal returns a dynamodb representation of the FuelStation struct.
func (site FuelStation) Marshal() (map[string]*dynamodb.AttributeValue, error) {
	fuelIds := []*dynamodb.AttributeValue{}
	fuelTypes := map[string]*dynamodb.AttributeValue{}
	for fuelId, price := range site.FuelTypes {
		fuelIds = append(fuelIds, &dynamodb.AttributeValue{
			N: aws.String(fmt.Sprintf("%d", fuelId)),
		})

		marshalledPrice, err := price.Marshal()
		if err != nil {
			return nil, err
		}

		fuelTypes[fmt.Sprintf("%d", fuelId)] = &dynamodb.AttributeValue{
			M: marshalledPrice,
		}
	}

	item := map[string]*dynamodb.AttributeValue{
		"SiteId": {
			N: aws.String(fmt.Sprintf("%d", site.SiteID)),
		},
		"LastUpdated": {
			S: aws.String(time.Now().Format(time.RFC3339)),
		},
		"FuelIds": {
			L: fuelIds,
		},
		"FuelTypes": {
			M: fuelTypes,
		},
	}
	return item, nil
}

// FuelStation.Marshal returns a dynamodb representation of the FuelStation struct.
func (site *FuelStation) Unmarshal(record map[string]*dynamodb.AttributeValue) error {
	siteIdRecord, ok := record["SiteId"]
	if !ok {
		return nil
	}
	siteId, err := strconv.Atoi(*siteIdRecord.N)
	if err != nil {
		return err
	}
	site.SiteID = siteId

	site.LastUpdated = SafeS(record, "LastUpdated")

	fuelTypesRecord, ok := record["FuelTypes"]
	if !ok {
		return nil
	}
	site.FuelTypes = map[int]FuelPrice{}
	for fuelIdRecord, fuelRecord := range fuelTypesRecord.M {
		fuelId, err := strconv.Atoi(fuelIdRecord)
		if err != nil {
			return err
		}

		var fuelPrice FuelPrice
		err = fuelPrice.Unmarshal(fuelRecord.M)
		if err != nil {
			return err
		}

		site.FuelTypes[fuelId] = fuelPrice
	}

	return nil
}

// FuelPriceList.Marshal returns a dynamodb representation of the FuelPriceList struct.
func (prices FuelPriceList) Marshal() ([]map[string]*dynamodb.AttributeValue, error) {
	items := []map[string]*dynamodb.AttributeValue{}
	for _, site := range prices.Sites {
		item, err := site.Marshal()
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (p *FuelPriceList) Unmarshal(records []map[string]*dynamodb.AttributeValue) error {
	p = &FuelPriceList{
		Sites: map[int]FuelStation{},
	}
	return nil
}

func AttachCorsHeaders(request *events.APIGatewayProxyResponse) {
	origin := defaults.Origin
	if defaults.IsLocal {
		origin = "*"
	}

	headers := map[string]string{
		"Access-Control-Allow-Headers": "*",
		"Access-Control-Allow-Origin":  origin,
		"Access-Control-Allow-Methods": "OPTIONS,GET,POST",
	}

	if request == nil {
		return
	}

	if request.Headers == nil {
		request.Headers = map[string]string{}
	}

	for k, v := range headers {
		request.Headers[k] = v
	}
}

func HandleCors(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Access-Control-Allow-Headers": "*",
			"Access-Control-Allow-Origin":  defaults.Origin,
			"Access-Control-Allow-Methods": "OPTIONS,GET,POST",
		},
	}, nil
}

func RespondWithStdErr(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       err.Error(),
		StatusCode: http.StatusInternalServerError,
	}, err
}

func GetClient(region string) *dynamodb.DynamoDB {
	config := aws.NewConfig().WithRegion(region)
	if defaults.IsLocal {
		fmt.Println("Using local endpoint.")
		config = config.WithEndpoint("http://dynamodb-local:8000")
	}

	session, err := session.NewSession()
	if err != nil {
		return nil
	}

	return dynamodb.New(session, config)
}

func CheckTableExists(client *dynamodb.DynamoDB, tableName string) bool {
	awsTables, err := client.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return false
	}

	tables := []string{}
	for _, table := range awsTables.TableNames {
		tables = append(tables, *table)
	}

	return slices.Contains(tables, tableName)
}

func S(v string) *string {
	return aws.String(v)
}
func N(v int) *string {
	return S(fmt.Sprintf("%d", v))
}
func F(v float64) *string {
	return S(decimal.NewFromFloat(v).String())
}

func SafeS(attrmap map[string]*dynamodb.AttributeValue, key string) string {
	if v, ok := attrmap[key]; ok {
		return *v.S
	}
	return "empty"
}

// GenerateSecureToken will create a new random token for use.
//
// https://stackoverflow.com/questions/45267125/how-to-generate-unique-random-alphanumeric-tokens-in-golang
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
