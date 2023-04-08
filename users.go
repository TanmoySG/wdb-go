package wdbgo

import (
	"fmt"

	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/routes"
	requestModels "github.com/TanmoySG/wdb-go/requestModels"
	wdbModels "github.com/TanmoySG/wunderDB/model"
)

const (
	failedLogin     = false
	successfulLogin = true
)

func (wdb wdbClient) LoginUser(username, password string) (bool, error) {
	queryEndpoint := routes.LoginUser.Format(wdb.ConnectionURI).String()
	queryMethod := methods.LoginUser.String()

	_, queryResponse, err := wdb.QueryClient.Query(queryEndpoint, queryMethod, nil)
	if err != nil {
		return failedLogin, err
	}

	apiResponse, err := queryResponse.ApiResponse()
	if err != nil {
		return failedLogin, err
	}

	if apiResponse.IsSuccess() {
		return successfulLogin, nil
	}

	return failedLogin, fmt.Errorf(apiResponse.Error.Code)
}

func (wdb wdbClient) CreateUser(username, password string) error {
	queryEndpoint := routes.CreateUser.Format(wdb.ConnectionURI).String()
	queryMethod := methods.CreateUser.String()
	queryPayload := requestModels.CreateUser{
		Username: username,
		Password: password,
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

func (wdb wdbClient) GrantRole(username, role string, database string, collection ...string) error {
	var targetDatabase, targetCollection *string

	targetDatabase = &database
	targetCollection = &collection[0]

	queryEndpoint := routes.GrantRoles.Format(wdb.ConnectionURI).String()
	queryMethod := methods.GrantRoles.String()
	queryPayload := requestModels.GrantRoles{
		Username: username,
		Permission: wdbModels.Permissions{
			Role: wdbModels.Identifier(role),
			On: &wdbModels.Entities{
				Databases:   targetDatabase,
				Collections: targetCollection,
			},
		},
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
