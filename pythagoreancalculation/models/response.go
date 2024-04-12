package models

// DistanceTypes response struct with closest and farthest coordinates
type DistanceTypes struct {
	Closest  [][]Coordinates
	Farthest [][]Coordinates
}
