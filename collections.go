package wdbgo

import (
	"fmt"

	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/routes"
	"github.com/TanmoySG/wdb-go/models"
	"github.com/TanmoySG/wdb-go/schema"
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
