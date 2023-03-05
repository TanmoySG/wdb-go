package routes

import "fmt"

type route string

const basicEndpointFormat = "%s/api/%s"
const queryParamsFormat = "%s?%s"

var (
	ApiPing route = ""

	LoginUser  route = "users/login"
	CreateUser route = "users"
	GrantRoles route = "users/grant"

	CreateRole route = "roles"
	ListRoles  route = "roles"

	AddData    route = "databases/%s/collections/%s/data"
	ReadData   route = "databases/%s/collections/%s/data"
	UpdateData route = "databases/%s/collections/%s/data"
	DeleteData route = "databases/%s/collections/%s/data"

	CreateCollection route = "databases/%s/collections"
	FetchCollection  route = "databases/%s/collections/%s"
	DeleteCollection route = "databases/%s/collections/%s"

	CreateDatabase route = "databases"
	FetchDatabase  route = "databases/%s"
	DeleteDatabase route = "databases/%s"
)

func (r route) Format(baseUrl string, endpointArgs ...any) route {
	endpointWithArgs := fmt.Sprintf(string(r), endpointArgs...)
	return route(fmt.Sprintf(basicEndpointFormat, baseUrl, endpointWithArgs))
}

func (r route) AddQueryParams(format string, queryParams ...any) route {
	query := fmt.Sprintf(format, queryParams...)
	endpointWithQueryParams := fmt.Sprintf(queryParamsFormat, r.String(), query)
	return route(endpointWithQueryParams)
}

func (r route) String() string {
	return string(r)
}
