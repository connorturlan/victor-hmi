package route

import (
	"fuelpriceservice/pkg/common"

	"github.com/aws/aws-lambda-go/events"
)

func GetRouteBetweenPoints(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// convert the json body to points
	// get the shortest route from the route api service
	// filter the stations based on the user's view
	return common.RespondWithStdErr(nil)
}
