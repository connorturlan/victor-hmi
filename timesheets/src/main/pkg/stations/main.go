package stations

import (
	"encoding/json"
	"errors"
	"fmt"
	"fuelpriceservice/pkg/common"
	"fuelpriceservice/pkg/common/defaults"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// GetStationPrices returns the prices that a given station has.
func GetStationPrices(client *dynamodb.DynamoDB, SiteID int) (common.FuelStation, error) {
	// verify the table exists.
	if !common.CheckTableExists(client, defaults.PricesTableName) {
		return common.FuelStation{}, errors.New("the database entry does not exist")
	}

	// get the station from the database.
	item := map[string]*dynamodb.AttributeValue{
		"SiteId": {N: aws.String(fmt.Sprintf("%d", SiteID))},
	}
	getRequest := &dynamodb.GetItemInput{
		TableName: aws.String(defaults.PricesTableName),
		Key:       item,
	}
	res, err := client.GetItem(getRequest)
	if err != nil {
		fmt.Printf("error while getting data from database. %s\n", err)
		return common.FuelStation{}, err
	}

	// unmarshal the dynamodb type
	data := common.FuelStation{}
	err = data.Unmarshal(res.Item)
	if err != nil {
		fmt.Printf("error while unmarshalling data from database. %s\n", err)
		return common.FuelStation{}, err
	}

	return data, nil
}

// GetStationDetails returns in depth information about a station.
func GetStationDetails(client *dynamodb.DynamoDB, SiteID int) (common.FuelStationDetails, error) {
	// verify the table exists.
	if !common.CheckTableExists(client, defaults.SitesTableName) {
		fmt.Printf("sites table doesn't exist.\n")
		return common.FuelStationDetails{}, errors.New("the database entry does not exist")
	}

	// get the station from the database.
	item := map[string]*dynamodb.AttributeValue{
		"SiteId": {N: aws.String(fmt.Sprintf("%d", SiteID))},
	}
	getRequest := &dynamodb.GetItemInput{
		TableName: aws.String(defaults.SitesTableName),
		Key:       item,
	}
	res, err := client.GetItem(getRequest)
	if err != nil {
		fmt.Printf("error while getting data from database. %s\n", err)
		return common.FuelStationDetails{}, err
	}

	// unmarshal the dynamodb type
	data := common.FuelStationDetails{}
	err = data.Unmarshal(res.Item)
	if err != nil {
		fmt.Printf("error while unmarshalling data from database. %s\n", err)
		return common.FuelStationDetails{}, err
	}

	return data, nil
}

// stations.HandleGet will query the database about a given siteid and return the full details for it.
func HandleGet(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("station get")

	// get client.
	client := common.GetClient(defaults.Region)

	// verify the table exists.
	if !common.CheckTableExists(client, defaults.PricesTableName) {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, errors.New("the database entry does not exist")
	}

	// get the siteid.
	SiteIDString, ok := request.QueryStringParameters["siteid"]
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, errors.New("the request is missing the siteid param")
	}
	SiteID, err := strconv.Atoi(SiteIDString)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, err
	}

	// get the price details.
	prices, err := GetStationPrices(client, SiteID)
	if err != nil {
		return common.RespondWithStdErr(err)
	}

	// get the station details.
	details, err := GetStationDetails(client, SiteID)
	if err != nil {
		return common.RespondWithStdErr(err)
	}
	details.LastUpdated = prices.LastUpdated
	details.FuelTypes = prices.FuelTypes

	// convert to json
	bytes, err := json.Marshal(details)
	if err != nil {
		fmt.Printf("error while marshalling data for response. %s\n", err)
		return common.RespondWithStdErr(err)
	}

	// return the body
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(bytes),
	}, nil
}
