package wdbgo

import (
	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/response"
	"github.com/TanmoySG/wdb-go/internal/routes"
)

func (wdb wdbClient) LoginUser(username, password string) (*apiResponse.Response, error) {
	queryEndpoint := routes.LoginUser.Format(wdb.ConnectionURI)
	queryMethod := methods.LoginUser.String()

	queryResponse, err := wdb.QueryClient.Query(queryEndpoint, queryMethod, nil)
	if err != nil {
		return nil, err
	}

	return queryResponse.ApiResponse()
}
