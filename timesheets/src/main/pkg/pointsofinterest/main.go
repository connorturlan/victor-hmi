package pointsofinterest

import (
	"encoding/json"
	"fmt"
	"fuelpriceservice/pkg/auth"
	"fuelpriceservice/pkg/common"
	"fuelpriceservice/pkg/common/defaults"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func GetUserData(client *dynamodb.DynamoDB, userid string) (auth.UserData, bool) {
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
		return auth.UserData{}, false
	}

	// get the saved maybeUser.
	maybeUser, ok := res.Item["UserData"]
	if !ok {
		fmt.Printf("user %s does not exist.\n", userid)
		return auth.UserData{}, false
	}
	userDataString := *maybeUser.S

	userData := auth.UserData{}
	err = json.Unmarshal([]byte(userDataString), &userData)
	if err != nil {
		fmt.Printf("[WARN] %s is an invalid userdata string, err: %s", userDataString, err)
		return auth.UserData{}, false
	}

	return userData, true
}

func HandleGet(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("poi get")

	// get the userid and token from the params, check the user is authorized.
	userid := auth.GetUserIdFromRequest(request)
	token := auth.GetTokenFromRequest(request)
	if !auth.IsTokenValid(userid, token) {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusForbidden,
			Body:       "{}",
		}, nil
	}

	// get client
	client := common.GetClient(defaults.Region)

	// verify the table exists
	if !common.CheckTableExists(client, defaults.UsersTableName) {
		// return the body
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	// get the user data.
	userData, okay := GetUserData(client, userid)
	if !okay {
		// return the body
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusForbidden,
		}, nil
	}

	// marshal the map
	bytes, err := json.Marshal(userData.PointsOfInterest)
	if err != nil {
		return common.RespondWithStdErr(err)
	}

	// return the body
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(bytes),
	}, nil
}

func HandlePost(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("poi set")

	// get the userid from the params.
	userid := auth.GetUserIdFromRequest(request)
	token := auth.GetTokenFromRequest(request)
	if !auth.IsTokenValid(userid, token) {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusForbidden,
		}, nil
	}

	// get client
	client := common.GetClient(defaults.Region)

	// verify the table exists
	if !common.CheckTableExists(client, defaults.UsersTableName) {
		// return the body
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	// get the user data.
	fmt.Println("getting existing user")
	userData, okay := GetUserData(client, userid)
	if !okay {
		// return the body
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusForbidden,
		}, nil
	}

	// create a points object.
	fmt.Println("inserting POIs")
	points := map[string]auth.Landmark{}
	err := json.Unmarshal([]byte(request.Body), &points)
	if err != nil {
		fmt.Printf("Error while unmarshalling userdata: %s\n", err)
		return common.RespondWithStdErr(err)
	}

	// insert the userdata.
	userData.PointsOfInterest = points
	bytes, err := json.Marshal(userData)
	if err != nil {
		fmt.Printf("Error while marshalling userdata: %s\n", err)
		return common.RespondWithStdErr(err)
	}

	// insert the user
	fmt.Printf("updating user. data: %s", string(bytes))
	item := map[string]*dynamodb.AttributeValue{
		"UserID":   {S: common.S(userid)},
		"UserData": {S: common.S(string(bytes))},
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
