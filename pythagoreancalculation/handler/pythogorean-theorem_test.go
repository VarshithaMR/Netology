package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"Netology/pythagoreancalculation/models"
)

func TestStartCalculateDistances(test *testing.T) {
	var testCases = []struct {
		testCaseName string
		jsonFile     string
		statusCode   int
	}{
		// Invalid in duplicating the testcase
		/*{
			testCaseName: "Invalid json request",
			jsonFile:     "./testutil/invalid-request.json",
			statusCode:   400,
		},*/
		{
			testCaseName: "no coordinates json request",
			jsonFile:     "./testutil/no-coordinates.json",
			statusCode:   400,
		},
		{
			testCaseName: "single coordinate json request",
			jsonFile:     "./testutil/single-coordinate.json",
			statusCode:   400,
		},
		{
			testCaseName: "Valid json request",
			jsonFile:     "./testutil/valid-request.json",
			statusCode:   200,
		},
	}

	for _, tc := range testCases {
		tc := tc
		test.Run(tc.testCaseName, func(t *testing.T) {
			t.Parallel()

			// Open and read the JSON file
			file, err := os.Open(tc.jsonFile)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			// read test json file
			requestData, err := io.ReadAll(file)
			if err != nil {
				test.Fatal(err)
			}

			//unmarshal test json file
			var requestBody []models.Coordinates
			if err := json.Unmarshal(requestData, &requestBody); err != nil {
				test.Fatal(err)
			}

			// Create a mock HTTP request
			req, err := http.NewRequest("POST", "/calculate", nil)
			if err != nil {
				test.Fatal(err)
			}
			req.Body = io.NopCloser(bytes.NewBuffer(requestData))

			// Create a mock HTTP response recorder
			resRecorder := httptest.NewRecorder()

			c := &calculateDomainHandlerService{}

			// Call the function with the mock request and response recorder
			c.StartCalculateDistances(req, resRecorder)

			// Check the HTTP response status code
			assert.Equal(test, tc.statusCode, resRecorder.Code)

		})
	}
}
