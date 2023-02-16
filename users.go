package wdbgo

import (
	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/response"
	"github.com/TanmoySG/wdb-go/internal/routes"
)

func (wdb wdbClient) LoginUser(username, password string) (*int, *apiResponse.Response, error) {
	queryEndpoint := routes.LoginUser.Format(wdb.ConnectionURI)
	queryMethod := methods.LoginUser.String()

	queryResponseStatus, queryResponse, err := wdb.QueryClient.Query(queryEndpoint, queryMethod, nil)
	if err != nil {
		return nil, nil, err
	}

	apiResponse, err := queryResponse.ApiResponse()
	if err != nil {
		return nil, nil, err
	}

	return queryResponseStatus, apiResponse, nil
}
