package auth

import (
	"fuelpriceservice/pkg/common"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (token *UserToken) Marshal() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"UserID":    {S: common.S(token.UserID)},
		"IssueTime": {S: common.S(token.IssueTime)},
		"Token":     {S: common.S(token.Token)},
	}
}

func (token *UserToken) Unmarshal(attrs map[string]*dynamodb.AttributeValue) {
	maybeUserID, ok := attrs["UserID"]
	if ok {
		token.UserID = (*maybeUserID.S)
	}
	maybeIssueTime, ok := attrs["IssueTime"]
	if ok {
		token.IssueTime = (*maybeIssueTime.S)
	}
	maybeToken, ok := attrs["Token"]
	if ok {
		token.Token = (*maybeToken.S)
	}
}
