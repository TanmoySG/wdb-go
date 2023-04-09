package wdbgo

import (
	"encoding/json"
	"fmt"

	dataFilters "github.com/TanmoySG/wdb-go/filters"
	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/routes"
	wdbModels "github.com/TanmoySG/wunderDB/model"
)

const dataKeyValueFilterFormat = "key=%s&value=%v"

type dataRecords map[wdbModels.Identifier]*wdbModels.Datum
type minifiedDataRecords map[string]interface{}

func (wdb wdbClient) AddData(data any, databaseName, collectionName string) error {
	queryEndpoint := routes.AddData.Format(wdb.ConnectionURI, databaseName, collectionName).String()
	queryMethod := methods.AddData.String()
	queryPayload := data

	_, queryResponse, err := wdb.QueryClient.Query(queryEndpoint, queryMethod, queryPayload)
	if err != nil {
		return err
	}

	apiResponse, err := queryResponse.ApiResponse()
	if err != nil {
		return err
	}

	if apiResponse.IsSuccess() {
		return nil
	}

	return fmt.Errorf(apiResponse.Error.Code)
}

func (wdb wdbClient) ReadData(databaseName, collectionName string, filters ...dataFilters.Filter) (dataRecords, error) {
	queryEndpoint := routes.ReadData.Format(wdb.ConnectionURI, databaseName, collectionName)
	queryMethod := methods.ReadData.String()

	filter, _ := resolveFilters(filters...)

	if filter != nil {
		queryEndpoint = queryEndpoint.AddQueryParams(dataKeyValueFilterFormat, filter.Key, filter.Value)
	}

	_, queryResponse, err := wdb.QueryClient.Query(queryEndpoint.String(), queryMethod, nil)
	if err != nil {
		return nil, err
	}

	apiResponse, err := queryResponse.ApiResponse()
	if err != nil {
		return nil, err
	}

	if apiResponse.IsSuccess() {
		var data map[wdbModels.Identifier]*wdbModels.Datum

		dataBytes, err := apiResponse.MarshalData()
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(dataBytes, &data)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return nil, fmt.Errorf(apiResponse.Error.Code)
}

func (dr dataRecords) Map() (map[string]interface{}, error) {
	dataByteArray, err := json.Marshal(dr)
	if err != nil {
		return nil, err
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(dataByteArray, &dataMap)
	if err != nil {
		return nil, err
	}

	return dataMap, nil
}

func (dr dataRecords) ByteArray() ([]byte, error) {
	dataByteArray, err := json.Marshal(dr)
	if err != nil {
		return nil, err
	}

	return dataByteArray, nil
}

func (dr dataRecords) String() (string, error) {
	dataByteArray, err := dr.ByteArray()
	if err != nil {
		return "nil", err
	}

	return string(dataByteArray), nil
}

func (dr dataRecords) Minified() minifiedDataRecords {
	minifiedDataRecords := make(minifiedDataRecords)
	for id, datum := range dr {
		minifiedDataRecords[id.String()] = datum.Data
	}
	return minifiedDataRecords
}

func (mdr minifiedDataRecords) Map() (map[string]interface{}, error) {
	dataByteArray, err := json.Marshal(mdr)
	if err != nil {
		return nil, err
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(dataByteArray, &dataMap)
	if err != nil {
		return nil, err
	}

	return dataMap, nil
}

func (mdr minifiedDataRecords) ByteArray() ([]byte, error) {
	dataByteArray, err := json.Marshal(mdr)
	if err != nil {
		return nil, err
	}

	return dataByteArray, nil
}

func (mdr minifiedDataRecords) String() (string, error) {
	dataByteArray, err := mdr.ByteArray()
	if err != nil {
		return "nil", err
	}

	return string(dataByteArray), nil
}

func (wdb wdbClient) UpdateData(dataPatch any, databaseName, collectionName string, filters ...dataFilters.Filter) error {
	queryEndpoint := routes.UpdateData.Format(wdb.ConnectionURI, databaseName, collectionName)
	queryMethod := methods.UpdateData.String()
	queryPayload := dataPatch

	filter, err := resolveFilters(filters...)
	if err != nil {
		return err
	}

	if filter == nil {
		return fmt.Errorf("filters not valid")
	}

	queryEndpoint = queryEndpoint.AddQueryParams(dataKeyValueFilterFormat, filter.Key, filter.Value)
	_, queryResponse, err := wdb.QueryClient.Query(queryEndpoint.String(), queryMethod, queryPayload)
	if err != nil {
		return err
	}

	apiResponse, err := queryResponse.ApiResponse()
	if err != nil {
		return err
	}

	if apiResponse.IsSuccess() {
		return nil
	}

	return fmt.Errorf(apiResponse.Error.Code)
}

func (wdb wdbClient) DeleteData(databaseName, collectionName string, filters ...dataFilters.Filter) error {
	queryEndpoint := routes.DeleteData.Format(wdb.ConnectionURI, databaseName, collectionName)
	queryMethod := methods.DeleteData.String()

	filter, err := resolveFilters(filters...)
	if err != nil {
		return err
	}

	if filter == nil {
		return fmt.Errorf("filters not valid")
	}

	queryEndpoint = queryEndpoint.AddQueryParams(dataKeyValueFilterFormat, filter.Key, filter.Value)
	_, queryResponse, err := wdb.QueryClient.Query(queryEndpoint.String(), queryMethod, nil)
	if err != nil {
		return err
	}

	apiResponse, err := queryResponse.ApiResponse()
	if err != nil {
		return err
	}

	if apiResponse.IsSuccess() {
		return nil
	}

	return fmt.Errorf(apiResponse.Error.Code)
}

func resolveFilters(filters ...dataFilters.Filter) (*dataFilters.Filter, error) {
	firstPickIndex := 0

	argsLen := len(filters)
	if argsLen == 0 {
		return nil, fmt.Errorf("filter missing")
	}

	// in case of multiple filters passed, pick first only
	pickedFilter := filters[firstPickIndex]
	if !pickedFilter.IsValid() {
		return nil, fmt.Errorf("invalid filter")
	}

	return &pickedFilter, nil
}
