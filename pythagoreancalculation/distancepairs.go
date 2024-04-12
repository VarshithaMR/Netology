package pythagoreancalculation

import (
	"errors"
	"sort"

	"Netology/pythagoreancalculation/models"
)

func MinMaxDistancePairs(pointsMap []models.PointsMap) (models.DistanceTypes, error) {
	if len(pointsMap) == 0 {
		return models.DistanceTypes{}, errors.New("empty Distance - coordinates list")
	}

	distancesList := make([]float64, len(pointsMap))
	for i, v := range pointsMap {
		distancesList[i] = v.Distance
	}

	sort.Float64s(distancesList)

	var (
		high, low [][]models.Coordinates
	)

	for i, v := range pointsMap {
		if v.Distance == distancesList[0] {
			low = pointsMap[i].CoordinatesList
		}
		if v.Distance == distancesList[len(distancesList)-1] {
			high = pointsMap[len(distancesList)-1].CoordinatesList
		}
		// Break the loop once both are found
		if len(low) > 0 && len(high) > 0 {
			break
		}
	}

	return models.DistanceTypes{
		Closest:  low,
		Farthest: high,
	}, nil
}
