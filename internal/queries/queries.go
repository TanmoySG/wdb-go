package queries

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	queryErrorFormat = "error processing query: %s"

	contentType = "application/json"

	contentTypeHeader = "Content-Type"

	userAgent = "User-Agent"
)

type queryResponse map[string]interface{}

type QueryClient struct {
	client         http.Client
	authentication authentication
	userAgent      string
}

type authentication struct {
	username string
	password string
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
	request.Header.Set(userAgent, qc.userAgent)
	
	// Content Type
	request.Header.Set(contentTypeHeader, contentType)

	response, err := qc.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf(queryErrorFormat, err)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(queryErrorFormat, err)
	}

	var queryResponse queryResponse
	err = json.Unmarshal([]byte(responseBody), &queryResponse)
	if err != nil {
		return nil, fmt.Errorf(queryErrorFormat, err)
	}

	return queryResponse, nil
}
