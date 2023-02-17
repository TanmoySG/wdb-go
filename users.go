package wdbgo

import (
	"fmt"

	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/routes"
	"github.com/TanmoySG/wdb-go/models"
)

const (
	failedLogin     = false
	successfulLogin = true
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

func (wdb wdbClient) CreateUser(username, password string) (bool, error) {
	queryEndpoint := routes.CreateUser.Format(wdb.ConnectionURI)
	queryMethod := methods.CreateUser.String()
	queryPayload := models.CreateUser{
		Username: username,
		Password: password,
	}

	_, queryResponse, err := wdb.QueryClient.Query(queryEndpoint, queryMethod, queryPayload)
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
