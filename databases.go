package wdbgo

import (
	"encoding/json"
	"fmt"

	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/routes"
	"github.com/TanmoySG/wdb-go/models"
	wdbModels "github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) CreateDatabase(databaseName string) error {
	queryEndpoint := routes.CreateDatabase.Format(wdb.ConnectionURI)
	queryMethod := methods.CreateDatabase.String()
	queryPayload := models.CreateDatabase{
		Name: databaseName,
	}

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

func (wdb wdbClient) GetDatabase(databaseName string) (*wdbModels.Database, error) {
	queryEndpoint := routes.FetchDatabase.Format(wdb.ConnectionURI, databaseName)
	queryMethod := methods.FetchDatabase.String()

	_, queryResponse, err := wdb.QueryClient.Query(queryEndpoint, queryMethod, nil)
	if err != nil {
		return nil, err
	}

	apiResponse, err := queryResponse.ApiResponse()
	if err != nil {
		return nil, err
	}

	if apiResponse.IsSuccess() {
		var database wdbModels.Database
		dataBytes, err := apiResponse.MarshalData()
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(dataBytes, &database)
		if err != nil {
			return nil, err
		}

		return &database, nil
	}

	return nil, fmt.Errorf(apiResponse.Error.Code)
}

func (wdb wdbClient) DeleteDatabase(databaseName string) error {
	queryEndpoint := routes.DeleteDatabase.Format(wdb.ConnectionURI, databaseName)
	queryMethod := methods.DeleteDatabase.String()

	_, queryResponse, err := wdb.QueryClient.Query(queryEndpoint, queryMethod, nil)
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
