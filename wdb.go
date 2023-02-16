package wdbgo

import (
	"fmt"

	"github.com/TanmoySG/wdb-go/internal/queries"
	"github.com/TanmoySG/wdb-go/internal/response"
	"github.com/TanmoySG/wdb-go/internal/version"
)

var (
	userAgent = fmt.Sprintf("wdb-go.client-library-%s", version.Version)
)

type Client interface {
	LoginUser(username, password string) (*int, *apiResponse.Response, error)
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

func NewWdbClient(username, password, ConnectionURI string, projectId *string) Client {
	ua := createUserAgent(projectId)
	return wdbClient{
		Username:      username,
		Password:      password,
		ConnectionURI: ConnectionURI,
		Metadata: wdbClientMetadata{
			UserAgent: ua,
		},
		QueryClient: queries.NewQueryClient(username, password, ua),
	}
}

func createUserAgent(projectId *string) string {
	if projectId != nil {
		return fmt.Sprintf("%s.projectId-%s", userAgent, *projectId)
	}
	return userAgent
}
