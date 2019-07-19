package smashteroids

import "github.com/faiface/pixel"

type Ship interface {
	Draw(t pixel.Target)
	Update(leftPressed, rightPressed, upPressed bool)
	ResetPosition()
	GetFuelUse() float64
}
