package history

import (
	"errors"
	"fuelpriceservice/pkg/common"
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type FuelData struct {
	Dates    []string      `json:"dates"`
	Datasets []FuelDataset `json:"datasets"`
}

type FuelDataset struct {
	Name   string    `json:"name"`
	FuelId int       `json:"fuelId"`
	Data   []float64 `json:"data"`
}

func (d FuelData) Marshal() (map[string]*dynamodb.AttributeValue, error) {
	dates := []*dynamodb.AttributeValue{}
	for _, date := range d.Dates {
		dates = append(dates, &dynamodb.AttributeValue{S: common.S(date)})
	}

	datasets := []*dynamodb.AttributeValue{}
	for _, set := range d.Datasets {
		marshalledSet, err := set.Marshal()
		if err != nil {
			return nil, err
		}

		datasets = append(datasets, &dynamodb.AttributeValue{M: marshalledSet})
	}

	return map[string]*dynamodb.AttributeValue{
		"dates": {
			L: dates,
		},
		"datasets": {
			L: datasets,
		},
	}, nil
}

func (d *FuelData) Unmarshal(values map[string]*dynamodb.AttributeValue) error {
	maybeDates, ok := values["dates"]
	if !ok {
		return errors.New("missing 'dates' from record")
	}
	dates := []string{}
	for _, dbvalues := range maybeDates.L {
		dates = append(dates, *dbvalues.S)
	}
	d.Dates = dates

	maybeData, ok := values["datasets"]
	if !ok {
		return errors.New("missing 'datasets' from record")
	}
	datasets := []FuelDataset{}
	for _, dbvalues := range maybeData.L {
		dataset := FuelDataset{}
		err := dataset.Unmarshal(dbvalues.M)
		if err != nil {
			return err
		}
		datasets = append(datasets, dataset)
	}
	d.Datasets = datasets

	return nil
}

func (d FuelDataset) Marshal() (map[string]*dynamodb.AttributeValue, error) {
	data := []*dynamodb.AttributeValue{}
	for _, price := range d.Data {
		data = append(data, &dynamodb.AttributeValue{N: common.F(price)})
	}

	return map[string]*dynamodb.AttributeValue{
		"name":   {S: common.S(d.Name)},
		"fuelid": {N: common.N(d.FuelId)},
		"data":   {L: data},
	}, nil
}

func (d *FuelDataset) Unmarshal(values map[string]*dynamodb.AttributeValue) error {
	maybeName, ok := values["name"]
	if !ok {
		return errors.New("missing 'name' from record")
	}
	d.Name = *maybeName.S

	maybeFuelId, ok := values["fuelid"]
	if !ok {
		return errors.New("missing 'fuelid' from record")
	}
	fuelId, err := strconv.Atoi(*maybeFuelId.N)
	if err != nil {
		return err
	}
	d.FuelId = fuelId

	maybeData, ok := values["data"]
	if !ok {
		return errors.New("missing 'data' from record")
	}
	data := []float64{}
	for _, dbvalue := range maybeData.L {
		value, err := strconv.ParseFloat(*dbvalue.N, 64)
		if err != nil {
			return err
		}
		data = append(data, value)
	}
	d.Data = data

	return nil
}
