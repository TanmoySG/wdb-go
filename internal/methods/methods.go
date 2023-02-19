package methods

import "net/http"

type method string

var (
	ApiPing method = http.MethodGet

	LoginUser  method = http.MethodGet
	CreateUser method = http.MethodPost
	GrantRoles method = http.MethodPost

	CreateRole method = http.MethodPost
	ListRoles  method = http.MethodGet

	AddData    method = http.MethodPost
	ReadData   method = http.MethodGet
	DeleteData method = http.MethodDelete
	UpdateData method = http.MethodPatch

	CreateCollection method = http.MethodPost
	FetchCollection  method = http.MethodGet
	DeleteCollection method = http.MethodDelete

	CreateDatabase method = http.MethodPost
	FetchDatabase  method = http.MethodGet
	DeleteDatabase method = http.MethodDelete
)

func (m method) String() string {
	return string(m)
}
