package auth

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

func GetUserIdFromRequest(request events.APIGatewayProxyRequest) string {

	userid, ok := request.QueryStringParameters["userid"]
	if !ok || userid == "" {
		fmt.Printf("no UserID provided.\n")
		return ""
	}
	return userid
}

func createUsersTable(client *dynamodb.DynamoDB) error {
	fmt.Println("Creating new users table!")

	_, err := client.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(defaults.UsersTableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("UserID"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("UserID"),
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

func HandleGet(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("auth login")

	// get client
	client := common.GetClient(defaults.Region)

	// verify the table exists
	if !common.CheckTableExists(client, defaults.UsersTableName) {
		createUsersTable(client)
	}

	// create a body object.
	userid, ok := request.QueryStringParameters["userid"]
	if !ok || userid == "" {
		fmt.Printf("no UserID provided.\n")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	// get the user from the database.
	item := map[string]*dynamodb.AttributeValue{
		"UserID": {S: common.S(userid)},
	}
	getRequest := &dynamodb.GetItemInput{
		TableName: aws.String(defaults.UsersTableName),
		Key:       item,
	}
	res, err := client.GetItem(getRequest)
	if err != nil {
		fmt.Printf("error while getting data from database. %s\n", err)
		return common.RespondWithStdErr(err)
	}

	// get the saved maybeUser.
	maybeUser, ok := res.Item["UserData"]
	if !ok {
		fmt.Printf("user %s does not exist.\n", userid)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusForbidden,
		}, nil
	}
	user := *maybeUser.S

	// check the item.
	fmt.Printf("user: %s\n", user)

	// return the body
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       user,
	}, nil
}

func HandlePost(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("auth post")

	// get client
	client := common.GetClient(defaults.Region)

	// verify the table exists
	if !common.CheckTableExists(client, defaults.UsersTableName) {
		createUsersTable(client)
	}

	// create a user object.
	user := UserData{}
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		fmt.Printf("Error while unmarshalling userdata: %s\n", err)
		return common.RespondWithStdErr(err)
	}

	// insert the user
	item := map[string]*dynamodb.AttributeValue{
		"UserID":   {S: common.S(user.UserID)},
		"UserData": {S: common.S(request.Body)},
	}
	req := &dynamodb.PutItemInput{
		TableName: aws.String(defaults.UsersTableName),
		Item:      item,
	}
	if _, err := client.PutItem(req); err != nil {
		fmt.Printf("error while writing data to database. %s\n", err)
		return common.RespondWithStdErr(err)
	}

	// return the body
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
	}, nil
}
