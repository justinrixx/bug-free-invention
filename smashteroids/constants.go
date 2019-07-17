package smashteroids

import "math"

const (
	radianScaleFactor = math.Pi / 180

	windowWidth  = 1024
	windowHeight = 764
)

func degreesToRadians(degrees int) float64 {
	return float64(degrees) * radianScaleFactor
}
