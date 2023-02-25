package wdbgo

import (
	"fmt"
	"net/http"

	"github.com/TanmoySG/wdb-go/internal/methods"
	"github.com/TanmoySG/wdb-go/internal/queries"
	"github.com/TanmoySG/wdb-go/internal/routes"
	"github.com/TanmoySG/wdb-go/internal/version"
	"github.com/TanmoySG/wdb-go/privileges"
	"github.com/TanmoySG/wdb-go/schema"
	wdbModels "github.com/TanmoySG/wunderDB/model"
)

const SkipConnectionCheck = false

var (
	userAgent = fmt.Sprintf("wdb-go.client-library-%s", version.Version)
)

type Client interface {
	Ping() (bool, error)

	LoginUser(username, password string) (bool, error)
	CreateUser(username, password string) error

	CreateRole(roleName string, allowedPrivileges, deniedPrivileges []privileges.Privilege) error
	GrantRoles(username, role string, entities ...string) error
	ListRoles() (map[string]wdbModels.Role, error)

	CreateDatabase(databaseName string) error
	GetDatabase(databaseName string) (*wdbModels.Database, error)
	DeleteDatabase(databaseName string) error

	CreateCollection(databaseName, collectionName string, schema schema.CollectionSchema) error
	GetCollection(databaseName, collectionName string) (*wdbModels.Collection, error)
	DeleteCollection(databaseName, collectionName string) error
}

type wdbClient struct {
	Username      string
	Password      string
	ConnectionURI string
	Metadata      wdbClientMetadata
	QueryClient   queries.QueryClient
}

type wdbClientMetadata struct {
	UserAgent string
}

func NewWdbClient(username, password, ConnectionURI string, projectId *string, args ...bool) (Client, error) {
	ua := createUserAgent(projectId)

	ok := testConnection(routes.ApiPing.Format(ConnectionURI).String(), args...)
	if !ok {
		return nil, fmt.Errorf("error creating wdb-client: connection failed")
	}

	return wdbClient{
		Username:      username,
		Password:      password,
		ConnectionURI: ConnectionURI,
		Metadata: wdbClientMetadata{
			UserAgent: ua,
		},
		QueryClient: queries.NewQueryClient(username, password, ua),
	}, nil
}

func createUserAgent(projectId *string) string {
	if projectId != nil {
		return fmt.Sprintf("%s.projectId-%s", userAgent, *projectId)
	}
	return userAgent
}

func (wdb wdbClient) Ping() (bool, error) {
	queryEndpoint := routes.ApiPing.Format(wdb.ConnectionURI).String()
	queryMethod := methods.ApiPing.String()

	_, queryResponse, err := wdb.QueryClient.Query(queryEndpoint, queryMethod, nil)
	if err != nil {
		return false, err
	}

	apiResponse, err := queryResponse.ApiResponse()
	if err != nil {
		return false, err
	}

	if apiResponse.IsSuccess() {
		return true, nil
	}

	return false, fmt.Errorf(apiResponse.Error.Code)
}

func testConnection(url string, args ...bool) bool {
	shouldCheck := !SkipConnectionCheck
	
	if len(args) != 0 {
		shouldCheck = args[0]
	}

	if shouldCheck {
		resp, err := http.Get(url)
		if err != nil {
			return err == nil
		}
		return resp.StatusCode == http.StatusOK
	}

	return true
}
