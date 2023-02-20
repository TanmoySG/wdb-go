package wdbgo

import (
	"fmt"

	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/routes"
	"github.com/TanmoySG/wdb-go/models"
	"github.com/TanmoySG/wunderDB/model"
)

const (
	failedLogin     = false
	successfulLogin = true

	failedCreateUser     = false
	successfulCreateUser = true
)

func (wdb wdbClient) LoginUser(username, password string) (bool, error) {
	queryEndpoint := routes.LoginUser.Format(wdb.ConnectionURI)
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
	queryEndpoint := routes.CreateUser.Format(wdb.ConnectionURI)
	queryMethod := methods.CreateUser.String()
	queryPayload := models.CreateUser{
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

func (wdb wdbClient) GrantRoles(username, role string, entities ...string) error {
	var targetDatabase, targetCollection *string

	entitiesCount := len(entities)
	if entitiesCount == 1 {
		targetDatabase = &entities[0]
	} else if entitiesCount == 2 {
		targetCollection = &entities[1]
	} else {
		return fmt.Errorf("entities missing: database or collection")
	}

	queryEndpoint := routes.GrantRoles.Format(wdb.ConnectionURI)
	queryMethod := methods.GrantRoles.String()
	queryPayload := models.GrantRoles{
		Username: username,
		Permission: model.Permissions{
			Role: model.Identifier(role),
			On: &model.Entities{
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
