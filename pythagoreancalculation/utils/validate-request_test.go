package utils

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"Netology/pythagoreancalculation/models"
)

func TestGetRequestBody(test *testing.T) {
	var testCases = []struct {
		testCaseName string
		input        io.ReadCloser
		response     []models.Coordinates
		err          error
	}{
		{
			testCaseName: "Valid Request format",
			input:        strings.NewReader(`[{"X": 10.5, "Y": 20.3}]`),
			response: []models.Coordinates{
				{
					X: 10.5,
					Y: 20.3,
				},
			},
			err: nil,
		},

		{
			testCaseName: "Invalid Request format",
			input:        strings.NewReader(`{}`),
			response:     nil,
			err:          &json.UnmarshalTypeError{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		test.Run(tc.testCaseName, func(t *testing.T) {
			t.Parallel()
			response, err := GetRequestBody(tc.input)
			if err != nil {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tc.response, response)
		})
	}
}
