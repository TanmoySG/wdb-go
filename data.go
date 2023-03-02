package wdbgo

import (
	"fmt"

	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/routes"
)

func (wdb wdbClient) AddData(data any, databaseName, collectionName string, args ...interface{}) error {
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
