package smashteroids

import (
	"math/rand"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

func GenerateTerrain(points int, initialHeight float64, bumpiness, roughness float64) []float64 {
	terrain := make([]float64, points)

	terrain[0] = initialHeight/2.0 + (rand.Float64() * bumpiness * 2) - bumpiness
	terrain[points-1] = terrain[0]
	bumpiness *= roughness

	for i := 1; i < points; i *= 2 {
		for j := (points / i) / 2; j < points-1; j += points / i {
			terrain[j] = ((terrain[j-(points/i)/2] + terrain[intMin(j+(points/i)/2, points-1)]) / 2)
			terrain[j] += (rand.Float64() * bumpiness * 2) - bumpiness
		}

		bumpiness *= roughness
	}

	return terrain
}

func DrawTerrain(terrain []float64, t pixel.Target, width float64) {
	imd := imdraw.New(nil)

	imd.Color = colornames.White

	delta := width / float64(len(terrain))
	currentPosition := 0.0
	for _, height := range terrain {
		imd.Push(pixel.V(currentPosition, height))
		currentPosition += delta
	}

	imd.Push(pixel.V(width+1, terrain[len(terrain)-1]))

	imd.Line(1)

	imd.Draw(t)
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
