package pythagoreancalculation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"Netology/pythagoreancalculation/models"
)

func TestCalculateDistance(test *testing.T) {
	var testcases = []struct {
		testCaseName string
		coordinates  []models.Coordinates
		response     models.DistanceTypes
		err          error
	}{
		{
			testCaseName: "Empty coordinates",
			coordinates:  []models.Coordinates{},
			response:     models.DistanceTypes{},
			err:          errors.New("less than two coordinates"),
		},
		{
			testCaseName: "single coordinate",
			coordinates: []models.Coordinates{
				{
					X: 1,
					Y: 4,
				},
			},
			response: models.DistanceTypes{},
			err:      errors.New("less than two coordinates"),
		},
		{
			testCaseName: "Multiple coordinate",
			coordinates: []models.Coordinates{
				{
					X: 1,
					Y: 2,
				},
				{
					X: 2,
					Y: 3,
				},
				{
					X: 3,
					Y: 4,
				},
			},
			response: models.DistanceTypes{
				Closest: [][]models.Coordinates{
					{
						{
							X: 1,
							Y: 2,
						},
						{
							X: 2,
							Y: 3,
						},
					},
					{
						{
							X: 2,
							Y: 3,
						},
						{
							X: 3,
							Y: 4,
						},
					},
				},
				Farthest: [][]models.Coordinates{
					{
						{
							X: 1,
							Y: 2,
						},
						{
							X: 3,
							Y: 4,
						},
					},
				},
			},
			err: nil,
		},
	}
	for _, tc := range testcases {
		tc := tc
		test.Run(tc.testCaseName, func(t *testing.T) {
			t.Parallel()
			response, err := CalculateDistance(tc.coordinates)
			if err != nil {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tc.response, response)
		})
	}
}
