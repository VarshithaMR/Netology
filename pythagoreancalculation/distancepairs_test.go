package pythagoreancalculation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"Netology/pythagoreancalculation/models"
)

func TestMinMaxDistancePairs(test *testing.T) {
	var testcases = []struct {
		testCaseName string
		request      []models.PointsMap
		response     models.DistanceTypes
		err          error
	}{
		{

			testCaseName: "Empty pointsmap",
			request:      []models.PointsMap{},
			response:     models.DistanceTypes{},
			err:          errors.New("empty Distance - coordinates list"),
		},
		{

			testCaseName: "Single request",
			request: []models.PointsMap{
				{
					Distance: 2.23,
					CoordinatesList: [][]models.Coordinates{
						{
							{
								X: 1,
								Y: 4,
							},
							{
								X: 3,
								Y: 2,
							},
						},
					},
				},
			},
			response: models.DistanceTypes{
				Closest: [][]models.Coordinates{
					{
						{
							X: 1,
							Y: 4,
						},
						{
							X: 3,
							Y: 2,
						},
					},
				},
				Farthest: [][]models.Coordinates{
					{
						{
							X: 1,
							Y: 4,
						},
						{
							X: 3,
							Y: 2,
						},
					},
				},
			},
			err: nil,
		},
		{

			testCaseName: "same distance multiple coordinates request",
			request: []models.PointsMap{
				{
					Distance: 2.23,
					CoordinatesList: [][]models.Coordinates{
						{
							{
								X: 1,
								Y: 4,
							},
							{
								X: 3,
								Y: 2,
							},
						},
						{
							{
								X: 4,
								Y: 4,
							},
							{
								X: 3,
								Y: 2,
							},
						},
					},
				},
				{
					Distance: 3,
					CoordinatesList: [][]models.Coordinates{
						{
							{
								X: 1,
								Y: 4,
							},
							{
								X: 4,
								Y: 4,
							},
						},
					},
				},
			},
			response: models.DistanceTypes{
				Closest: [][]models.Coordinates{
					{
						{
							X: 1,
							Y: 4,
						},
						{
							X: 3,
							Y: 2,
						},
					},
					{
						{
							X: 4,
							Y: 4,
						},
						{
							X: 3,
							Y: 2,
						},
					},
				},
				Farthest: [][]models.Coordinates{
					{
						{
							X: 1,
							Y: 4,
						},
						{
							X: 4,
							Y: 4,
						},
					},
				},
			},
			err: nil,
		},
		{

			testCaseName: "Multiple request",
			request: []models.PointsMap{
				{
					Distance: 2.23,
					CoordinatesList: [][]models.Coordinates{
						{
							{
								X: 3,
								Y: 2,
							},
							{
								X: 5,
								Y: 1,
							},
						},
					},
				},

				{
					Distance: 3,
					CoordinatesList: [][]models.Coordinates{
						{
							{
								X: 1,
								Y: 4,
							},
							{
								X: 4,
								Y: 4,
							},
						},
					},
				},
			},
			response: models.DistanceTypes{
				Closest: [][]models.Coordinates{
					{
						{
							X: 3,
							Y: 2,
						},
						{
							X: 5,
							Y: 1,
						},
					},
				},
				Farthest: [][]models.Coordinates{
					{
						{
							X: 1,
							Y: 4,
						},
						{
							X: 4,
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
			response, err := MinMaxDistancePairs(tc.request)
			if err != nil {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tc.response, response)
		})
	}
}
