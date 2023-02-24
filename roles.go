package wdbgo

import (
	"encoding/json"
	"fmt"

	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/routes"
	"github.com/TanmoySG/wdb-go/models"
	"github.com/TanmoySG/wdb-go/privileges"
	wdbModels "github.com/TanmoySG/wunderDB/model"
)

func (wdb wdbClient) CreateRole(roleName string, allowedPrivileges, deniedPrivileges []privileges.Privilege) error {
	var allowed, denied []string

	for _, allowedPrivilege := range allowedPrivileges {
		allowed = append(allowed, allowedPrivilege.Name())
	}

	for _, deniedPrivilege := range deniedPrivileges {
		denied = append(denied, deniedPrivilege.Name())
	}

	queryEndpoint := routes.CreateRole.Format(wdb.ConnectionURI)
	queryMethod := methods.CreateRole.String()
	queryPayload := models.CreateRole{
		Role:    roleName,
		Allowed: allowed,
		Denied:  denied,
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

func (wdb wdbClient) ListRoles() (map[string]wdbModels.Role, error) {

	queryEndpoint := routes.ListRoles.Format(wdb.ConnectionURI)
	queryMethod := methods.ListRoles.String()

	_, queryResponse, err := wdb.QueryClient.Query(queryEndpoint, queryMethod, nil)
	if err != nil {
		return nil, err
	}

	apiResponse, err := queryResponse.ApiResponse()
	if err != nil {
		return nil, err
	}

	if apiResponse.IsSuccess() {
		var rolesList map[string]wdbModels.Role
		dataBytes, err := apiResponse.MarshalData()
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(dataBytes, &rolesList)
		if err != nil {
			return nil, err
		}

		return rolesList, nil
	}

	return nil, fmt.Errorf(apiResponse.Error.Code)
}
