package history

import (
	"encoding/json"
	"fmt"
	"fuelpriceservice/pkg/common"
	"fuelpriceservice/pkg/common/defaults"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func createHistoryTable(client *dynamodb.DynamoDB) error {
	fmt.Println("Creating new history table!")

	_, err := client.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(defaults.HistoryTableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Filename"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Filename"),
				KeyType:       aws.String("HASH"),
			},
		},

		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	})

	return err
}

func HandlePost(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("history post")
	// get client
	client := common.GetClient(defaults.Region)

	// verify the table exists
	if !common.CheckTableExists(client, defaults.HistoryTableName) {
		createHistoryTable(client)
	}

	// unmarshal the body
	data := FuelData{}
	err := json.Unmarshal([]byte(request.Body), &data)
	if err != nil {
		fmt.Printf("error while unmarshalling request body. %s\n", err)
		return common.RespondWithStdErr(err)
	}

	// convert to dynamodb type
	values, err := data.Marshal()
	if err != nil {
		fmt.Printf("error while marshalling data for database. %s\n", err)
		return common.RespondWithStdErr(err)
	}
	values["Filename"] = &dynamodb.AttributeValue{S: common.S("history")}

	// post to table
	req := &dynamodb.PutItemInput{
		TableName: aws.String(defaults.HistoryTableName),
		Item:      values,
	}
	if _, err := client.PutItem(req); err != nil {
		fmt.Printf("error while writing data to database. %s\n", err)
		return common.RespondWithStdErr(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
	}, nil
}

func HandleGet(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("history get")

	// get client
	client := common.GetClient(defaults.Region)

	// verify the table exists
	if !common.CheckTableExists(client, defaults.HistoryTableName) {
		createHistoryTable(client)
	}

	// get the file from the database.
	values := map[string]*dynamodb.AttributeValue{
		"Filename": {S: common.S("history")},
	}
	getRequest := &dynamodb.GetItemInput{
		TableName: aws.String(defaults.HistoryTableName),
		Key:       values,
	}
	res, err := client.GetItem(getRequest)
	if err != nil {
		fmt.Printf("error while getting data from database. %s\n", err)
		return common.RespondWithStdErr(err)
	}

	// unmarshal the dynamodb type
	data := FuelData{}
	err = data.Unmarshal(res.Item)
	if err != nil {
		fmt.Printf("error while unmarshalling data from database. %s\n", err)
		return common.RespondWithStdErr(err)
	}

	// convert to json
	bytes, err := json.Marshal(data)
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
