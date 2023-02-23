package wdbgo

import (
	"encoding/json"
	"fmt"

	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/routes"
	"github.com/TanmoySG/wdb-go/models"
	"github.com/TanmoySG/wdb-go/schema"

	wdbModels "github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) CreateCollection(databaseName, collectionName string, schema schema.CollectionSchema) error {
	queryEndpoint := routes.CreateCollection.Format(wdb.ConnectionURI, databaseName)
	queryMethod := methods.CreateCollection.String()
	queryPayload := models.CreateCollection{
		Name:   collectionName,
		Schema: schema,
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

func (wdb wdbClient) GetCollection(databaseName, collectionName string) (*wdbModels.Collection, error) {
	queryEndpoint := routes.FetchCollection.Format(wdb.ConnectionURI, databaseName, collectionName)
	queryMethod := methods.FetchCollection.String()

	_, queryResponse, err := wdb.QueryClient.Query(queryEndpoint, queryMethod, nil)
	if err != nil {
		return nil, err
	}

	apiResponse, err := queryResponse.ApiResponse()
	if err != nil {
		return nil, err
	}

	if apiResponse.IsSuccess() {
		var collection wdbModels.Collection
		dataBytes, err := apiResponse.MarshalData()
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(dataBytes, &collection)
		if err != nil {
			return nil, err
		}

		return &collection, nil
	}

	return nil, fmt.Errorf(apiResponse.Error.Code)
}
