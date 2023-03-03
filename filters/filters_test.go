package dataFilters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Filter_IsValid(t *testing.T) {

	type testCase struct {
		filter         Filter
		expectedResult bool
	}

	testCases := []testCase{
		{
			filter: Filter{
				Key:   "name",
				Value: "John",
			},
			expectedResult: true,
		},
		{
			filter: Filter{
				Value: "John",
			},
			expectedResult: false,
		},
		{
			filter: Filter{
				Key: "name",
			},
			expectedResult: false,
		},
		{
			filter:         Filter{},
			expectedResult: false,
		},
	}

	for _, tc := range testCases {

		isValidFilter := tc.filter.IsValid()
		assert.Equal(t, tc.expectedResult, isValidFilter)
	}
}
