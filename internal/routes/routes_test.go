package routes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseUrl = "abc.xyz"

var expectedDataEndpoints = fmt.Sprintf("databases/%s/collections/%s", "database", "collection")

func Test_Format(t *testing.T) {
	type testCase struct {
		baseUrl              string
		route                route
		endpointArgs         []interface{}
		expectedFormattedURL string
	}

	testCases := []testCase{
		{
			baseUrl:              baseUrl,
			route:                LoginUser,
			endpointArgs:         nil,
			expectedFormattedURL: baseUrl + "/api/users/login",
		},
		{
			baseUrl:              baseUrl,
			route:                CreateUser,
			endpointArgs:         nil,
			expectedFormattedURL: baseUrl + "/api/users",
		},
		{
			baseUrl:              baseUrl,
			route:                GrantRoles,
			endpointArgs:         nil,
			expectedFormattedURL: baseUrl + "/api/users/grant",
		},
		{
			baseUrl:              baseUrl,
			route:                CreateRole,
			endpointArgs:         nil,
			expectedFormattedURL: baseUrl + "/api/roles",
		},
		{
			baseUrl:              baseUrl,
			route:                ListRoles,
			endpointArgs:         nil,
			expectedFormattedURL: baseUrl + "/api/roles",
		},
		{
			baseUrl:              baseUrl,
			route:                AddData,
			endpointArgs:         []interface{}{"database", "collection"},
			expectedFormattedURL: baseUrl + "/api/" + expectedDataEndpoints + "/data",
		},
		{
			baseUrl:              baseUrl,
			route:                ReadData,
			endpointArgs:         []interface{}{"database", "collection"},
			expectedFormattedURL: baseUrl + "/api/" + expectedDataEndpoints + "/data",
		},
		{
			baseUrl:              baseUrl,
			route:                DeleteData,
			endpointArgs:         []interface{}{"database", "collection"},
			expectedFormattedURL: baseUrl + "/api/" + expectedDataEndpoints + "/data",
		},
		{
			baseUrl:              baseUrl,
			route:                UpdateData,
			endpointArgs:         []interface{}{"database", "collection"},
			expectedFormattedURL: baseUrl + "/api/" + expectedDataEndpoints + "/data",
		},
		{
			baseUrl:              baseUrl,
			route:                CreateCollection,
			endpointArgs:         []interface{}{"database"},
			expectedFormattedURL: baseUrl + "/api/" + "databases/database/collections",
		},
		{
			baseUrl:              baseUrl,
			route:                FetchCollection,
			endpointArgs:         []interface{}{"database", "collection"},
			expectedFormattedURL: baseUrl + "/api/" + expectedDataEndpoints,
		},
		{
			baseUrl:              baseUrl,
			route:                DeleteCollection,
			endpointArgs:         []interface{}{"database", "collection"},
			expectedFormattedURL: baseUrl + "/api/" + expectedDataEndpoints,
		},
		{
			baseUrl:              baseUrl,
			route:                CreateDatabase,
			endpointArgs:         nil,
			expectedFormattedURL: baseUrl + "/api/databases",
		},
		{
			baseUrl:              baseUrl,
			route:                DeleteDatabase,
			endpointArgs:         []interface{}{"database"},
			expectedFormattedURL: baseUrl + "/api/databases/database",
		},
		{
			baseUrl:              baseUrl,
			route:                FetchDatabase,
			endpointArgs:         []interface{}{"database"},
			expectedFormattedURL: baseUrl + "/api/databases/database",
		},
	}

	for _, tc := range testCases {
		var formatResult string

		if len(tc.endpointArgs) == 0 || tc.endpointArgs == nil {
			formatResult = tc.route.Format(tc.baseUrl).String()
		} else {
			formatResult = tc.route.Format(tc.baseUrl, tc.endpointArgs...).String()
		}

		assert.Equal(t, tc.expectedFormattedURL, formatResult)
	}

}

func Test_AddQueryParams(t *testing.T) {
	type testCase struct {
		baseUrl                             string
		route                               route
		endpointArgs                        []interface{}
		expectedFormattedURL                string
		queryFormat                         string
		queryParams                         []interface{}
		expectedFormattedUrlWithQueryParams string
	}

	testCases := []testCase{
		{
			baseUrl:                             baseUrl,
			route:                               AddData,
			endpointArgs:                        []interface{}{"database", "collection"},
			expectedFormattedURL:                baseUrl + "/api/" + expectedDataEndpoints + "/data",
			queryFormat:                         "key=%s&value=%s",
			queryParams:                         []interface{}{"key", "value"},
			expectedFormattedUrlWithQueryParams: baseUrl + "/api/" + expectedDataEndpoints + "/data" + "?key=key&value=value",
		},
		{
			baseUrl:                             baseUrl,
			route:                               ReadData,
			endpointArgs:                        []interface{}{"database", "collection"},
			expectedFormattedURL:                baseUrl + "/api/" + expectedDataEndpoints + "/data",
			queryFormat:                         "key=%s&value=%s",
			queryParams:                         []interface{}{"key", "value"},
			expectedFormattedUrlWithQueryParams: baseUrl + "/api/" + expectedDataEndpoints + "/data" + "?key=key&value=value",
		},
		{
			baseUrl:                             baseUrl,
			route:                               DeleteData,
			endpointArgs:                        []interface{}{"database", "collection"},
			expectedFormattedURL:                baseUrl + "/api/" + expectedDataEndpoints + "/data",
			queryFormat:                         "key=%s&value=%s",
			queryParams:                         []interface{}{"key", "value"},
			expectedFormattedUrlWithQueryParams: baseUrl + "/api/" + expectedDataEndpoints + "/data" + "?key=key&value=value",
		},
		{
			baseUrl:                             baseUrl,
			route:                               UpdateData,
			endpointArgs:                        []interface{}{"database", "collection"},
			expectedFormattedURL:                baseUrl + "/api/" + expectedDataEndpoints + "/data",
			queryFormat:                         "key=%s&value=%s",
			queryParams:                         []interface{}{"key", "value"},
			expectedFormattedUrlWithQueryParams: baseUrl + "/api/" + expectedDataEndpoints + "/data" + "?key=key&value=value",
		},
	}

	for _, tc := range testCases {
		var formatResult, epWithQueryParam route

		if len(tc.endpointArgs) == 0 || tc.endpointArgs == nil {
			formatResult = tc.route.Format(tc.baseUrl)
		} else {
			formatResult = tc.route.Format(tc.baseUrl, tc.endpointArgs...)
		}

		epWithQueryParam = formatResult.AddQueryParams(tc.queryFormat, tc.queryParams...)

		assert.Equal(t, tc.expectedFormattedURL, formatResult.String())
		assert.Equal(t, tc.expectedFormattedUrlWithQueryParams, epWithQueryParam.String())

	}

}
