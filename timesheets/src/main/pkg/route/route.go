package route

import (
	"fuelpriceservice/pkg/common"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (p *Point) ToAttribute() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"lat": {S: common.S(p.Lat)},
		"lng": {S: common.S(p.Lng)},
	}
}

func (p *Point) FromAttribute(attrs map[string]*dynamodb.AttributeValue) {
	maybeLat, ok := attrs["lat"]
	if ok {
		p.Lat = (*maybeLat.S)
	}
	maybeLng, ok := attrs["lng"]
	if ok {
		p.Lng = (*maybeLng.S)
	}
}
