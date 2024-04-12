package pythagoreancalculation

import (
	"errors"
	"fmt"
	"log"
	"math"

	"Netology/pythagoreancalculation/models"
)

/*
CalculateDistance function to calculate the distance between 2 points using
Pythagorean theorem : d=√((x_2-x_1)²+(y_2-y_1)²)
*/
func CalculateDistance(points []models.Coordinates) (models.DistanceTypes, error) {
	if len(points) < 2 {
		return models.DistanceTypes{}, errors.New("less than two coordinates")
	}
	var (
		xCoordinates, yCoordinates, distance float64
		coordinatesWithDistance              []models.PointsMap
	)

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {

			x1, y1 := points[i].X, points[i].Y
			x2, y2 := points[j].X, points[j].Y

			log.Printf("Coordinate Pairs \n 1. %f : %f\n 2. %f : %f \n", x1, y1, x2, y2)

			xCoordinates = math.Pow(x2-x1, 2)
			yCoordinates = math.Pow(y2-y1, 2)

			distance = math.Sqrt(xCoordinates + yCoordinates)
			log.Printf("Distance between the coordinates = %f\n", distance)

			found := false
			// Check if the distance already exists in coordinatesWithDistance
			for i, entry := range coordinatesWithDistance {
				if entry.Distance == distance {
					// Update the coordinates list for this distance
					coordinatesWithDistance[i].CoordinatesList = append(coordinatesWithDistance[i].CoordinatesList,
						[]models.Coordinates{{X: x1, Y: y1}, {X: x2, Y: y2}})
					found = true
					break
				}
			}

			// If the distance doesn't exist, add a new entry to coordinatesWithDistance
			if !found {
				coordinatesWithDistance = append(coordinatesWithDistance, models.PointsMap{
					Distance:        distance,
					CoordinatesList: [][]models.Coordinates{{models.Coordinates{X: x1, Y: y1}, models.Coordinates{X: x2, Y: y2}}},
				})
			}
		}
		fmt.Println()
	}
	distancePairsResponse, err := MinMaxDistancePairs(coordinatesWithDistance)
	if err != nil {
		msg := "improper distance Pairs formed"
		log.Printf(msg)
		return distancePairsResponse, errors.New(msg)
	}

	return distancePairsResponse, nil

}
