package queries

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/TanmoySG/wdb-go/internal/response"
)

const (
	queryErrorFormat    = "error processing query: %s"
	responseErrorFormat = "error processing response: %s"

	contentType       = "application/json"
	contentTypeHeader = "Content-Type"

	userAgentHeader = "User-Agent"
)

type queryResponse []byte

type QueryClient struct {
	client         http.Client
	authentication authentication
	userAgent      string
}

type authentication struct {
	username string
	password string
}

func NewQueryClient(username, password, userAgent string) QueryClient {
	qc := QueryClient{
		client: *http.DefaultClient,
		authentication: authentication{
			username: username,
			password: password,
		},
		userAgent: userAgent,
	}

	return qc
}

func (qc QueryClient) Query(endpoint, method string, payload interface{}) (queryResponse, error) {

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf(queryErrorFormat, err)
	}

	request, err := http.NewRequest(method, endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf(queryErrorFormat, err)
	}

	// can change based on requirement - make it flexible
	request.SetBasicAuth(qc.authentication.username, qc.authentication.password)

	// User Agent set to wdb-go-client-<project name>
	request.Header.Set(userAgentHeader, qc.userAgent)

	// Content Type
	request.Header.Set(contentTypeHeader, contentType)

	response, err := qc.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf(queryErrorFormat, err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(responseErrorFormat, err)
	}

	return responseBody, nil
}

func (qr queryResponse) Map() (map[string]interface{}, error) {
	var responseMap map[string]interface{}
	err := json.Unmarshal(qr, &responseMap)
	if err != nil {
		return nil, fmt.Errorf(responseErrorFormat, err)
	}
	return responseMap, nil
}

func (qr queryResponse) ApiResponse() (*apiResponse.Response, error) {
	var apiResponse apiResponse.Response
	err := json.Unmarshal(qr, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf(responseErrorFormat, err)
	}
	return &apiResponse, nil
}
