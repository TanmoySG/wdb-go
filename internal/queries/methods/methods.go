package methods

import "net/http"

type method string

var (
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
