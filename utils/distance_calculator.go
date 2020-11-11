package utils

import "math"

const (
	earthRadiusMi = 3958 // radius of the earth in miles.
	earthRaidusKm = 6371 // radius of the earth in kilometers.
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

func DistanceCalculator(coordinatesA, coordinatesB Coordinates) (float64, float64) {
	lat1 := toRad(coordinatesA.Latitude)
	lon1 := toRad(coordinatesA.Longitude)
	lat2 := toRad(coordinatesB.Latitude)
	lon2 := toRad(coordinatesB.Longitude)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	mi := c * earthRadiusMi
	km := c * earthRaidusKm

	return mi, km
}

func toRad(value float64) float64 {
	return value * math.Pi / 180
}
