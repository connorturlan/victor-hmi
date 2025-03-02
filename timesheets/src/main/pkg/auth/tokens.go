package auth

import (
	"fmt"
	"fuelpriceservice/pkg/common"
	"fuelpriceservice/pkg/common/defaults"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	StatusOk uint = iota
	StatusErrorRequest
	StatusTokenInvalid
)

func createUserTokensTable(client *dynamodb.DynamoDB) error {
	fmt.Println("Creating new user tokens table!")

	_, err := client.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String(defaults.UserTokensTableName),
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

func IsTokenValid(userid, token string) bool {
	return CheckTokenValid(userid, token) == 0
}

func GetTokenFromRequest(request events.APIGatewayProxyRequest) string {
	authorisation, ok := request.Headers["Authorization"]
	if !ok || authorisation == "" {
		fmt.Printf("no UserToken provided.\n")
		return ""
	}
	splitAuth := strings.Split(authorisation, " ")
	token := ""
	if len(splitAuth) > 1 {
		token = splitAuth[1]
	}
	return token
}

// ProvideToken will generate a new token for a user, returned within a json body.
//
// It assumes the user has been authenticated, generates a new random token, and inserts it into the database.
func ProvideToken(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("ProvideToken")

	// get client
	client := common.GetClient(defaults.Region)

	// verify the table exists
	if !common.CheckTableExists(client, defaults.UserTokensTableName) {
		createUserTokensTable(client)
	}

	// create a body object.
	userid := GetUserIdFromRequest(request)

	// create a new token.
	token := UserToken{
		UserID:    userid,
		IssueTime: time.Now().Format(time.RFC3339),
		Token:     common.GenerateSecureToken(32),
	}

	// insert the token into the database.
	item := token.Marshal()
	req := &dynamodb.PutItemInput{
		TableName: aws.String(defaults.UserTokensTableName),
		Item:      item,
	}
	if _, err := client.PutItem(req); err != nil {
		fmt.Printf("error while writing data to database. %s\n", err)
		return common.RespondWithStdErr(err)
	}

	// return the body
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusAccepted,
		Body:       token.Token,
	}, nil
}

func CheckTokenValid(userid, token string) uint {
	// get client
	client := common.GetClient(defaults.Region)

	// verify the table exists
	if !common.CheckTableExists(client, defaults.UserTokensTableName) {
		createUserTokensTable(client)
	}

	// get the user from the database.
	item := map[string]*dynamodb.AttributeValue{
		"UserID": {S: common.S(userid)},
	}
	getRequest := &dynamodb.GetItemInput{
		TableName: aws.String(defaults.UserTokensTableName),
		Key:       item,
	}
	res, err := client.GetItem(getRequest)
	if err != nil {
		fmt.Printf("error while getting data from database. %s\n", err)
		return StatusErrorRequest
	}

	// check the tokens against one another.
	var user UserToken
	user.Unmarshal(res.Item)
	if user.Token == "" || user.Token != token {
		return StatusTokenInvalid
	}

	return StatusOk
}

func ValidateToken(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("ValidateToken")

	// get client
	client := common.GetClient(defaults.Region)

	// verify the table exists
	if !common.CheckTableExists(client, defaults.UserTokensTableName) {
		createUserTokensTable(client)
	}

	// create a body object.
	userid := GetUserIdFromRequest(request)
	token := GetTokenFromRequest(request)

	// check the token and return the appropriate response.
	switch CheckTokenValid(userid, token) {
	case StatusOk:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusAccepted,
		}, nil
	case StatusTokenInvalid:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusForbidden,
		}, nil
	case StatusErrorRequest:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	default:
		fmt.Printf("the check did not return the expected value.")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, nil
	}
}
