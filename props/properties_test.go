package props

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadProperties(test *testing.T) {
	var testcases = []struct {
		testCaseName   string
		propertiesFile string
		response       *Properties
		err            error
	}{
		{
			testCaseName:   "all properties",
			propertiesFile: "./testutil/application_test.yaml",
			response: &Properties{
				Server: ServerProps{
					Host:        "localhost",
					Port:        8080,
					ContextRoot: "/calculate",
				},
			},
			err: nil,
		},
		{
			testCaseName:   "missing properties",
			propertiesFile: "./testutil/application-noport_test.yaml",
			response: &Properties{
				Server: ServerProps{
					Host:        "",
					Port:        0,
					ContextRoot: "/calculate",
				},
			},
			err: nil,
		},
		{
			testCaseName:   "improper file",
			propertiesFile: "./testutil/application-improper_test.yaml",
			response: &Properties{
				Server: ServerProps{
					Host:        "",
					Port:        0,
					ContextRoot: "",
				},
			},
			err: nil,
		},
	}
	for _, tc := range testcases {
		tc := tc
		test.Run(tc.testCaseName, func(t *testing.T) {
			t.Parallel()
			response, err := ReadProperties(tc.propertiesFile)
			assert.Equal(t, tc.response, response)
			assert.Equal(t, tc.err, err)
		})
	}
}
