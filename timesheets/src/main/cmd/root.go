package cmd

import (
	"fmt"
	"fuelpriceservice/pkg/auth"
	"fuelpriceservice/pkg/common"
	"fuelpriceservice/pkg/history"
	"fuelpriceservice/pkg/pointsofinterest"
	"fuelpriceservice/pkg/stations"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(request events.APIGatewayProxyRequest) (res events.APIGatewayProxyResponse, err error) {
	res = events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Hello, World!",
	}
	return

	fmt.Printf("routing %s -- %s\n", request.HTTPMethod, request.Path)
	switch request.HTTPMethod {
	case http.MethodPost:
		fmt.Printf("routing %s -- %s in post\n", request.HTTPMethod, request.Path)
		switch request.Path {
		case "/average":
			res, err = history.HandlePost(request)
		case "/auth":
			res, err = auth.HandlePost(request)
		case "/poi":
			res, err = pointsofinterest.HandlePost(request)
		default:
			res = events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
			}
		}
	case http.MethodGet:
		fmt.Printf("routing %s -- %s in get\n", request.HTTPMethod, request.Path)
		switch request.Path {
		case "/history":
			res, err = history.HandleGet(request)
		case "/login":
			res, err = auth.HandleGet(request)
		case "/poi":
			res, err = pointsofinterest.HandleGet(request)
		case "/station":
			res, err = stations.HandleGet(request)
		case "/token":
			res, err = auth.ProvideToken(request)
		case "/validate":
			res, err = auth.ValidateToken(request)
		default:
			res = events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
			}
		}
	case http.MethodOptions:
		fmt.Printf("routing %s -- %s in options\n", request.HTTPMethod, request.Path)
		res, err = common.HandleCors(request)
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	// return an error if applicable.
	if err != nil {
		common.RespondWithStdErr(err)
	}

	common.AttachCorsHeaders(&res)
	return
}
