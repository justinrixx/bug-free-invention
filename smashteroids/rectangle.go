package smashteroids

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Rectangle struct {
	Rect   pixel.Rect
	Matrix pixel.Matrix
	Color  color.Color
}

func (r *Rectangle) Draw(imd *imdraw.IMDraw) {
	imd.Color = r.Color
	imd.Push(r.Rect.Min, r.Rect.Max)
	imd.Rectangle(0)
}
