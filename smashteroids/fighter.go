package smashteroids

import (
	"math"

	"github.com/faiface/pixel"
)

type Fighter struct {
	Position pixel.Vec
	Rotation int
	Sprite   *pixel.Sprite

	d pixel.Vec

	positionLeft  pixel.Vec
	positionRight pixel.Vec

	rotationRadians float64
}

func (_ *Fighter) GetFuelUse() float64 {
	return 1.0
}

func (f *Fighter) Draw(t pixel.Target) {
	mat := pixel.IM.Moved(f.Position).Rotated(f.Position, f.rotationRadians)
	matLeft := pixel.IM.Moved(f.positionLeft).Rotated(f.positionLeft, f.rotationRadians)
	matRight := pixel.IM.Moved(f.positionRight).Rotated(f.positionRight, f.rotationRadians)

	f.Sprite.Draw(t, mat)
	f.Sprite.Draw(t, matLeft)
	f.Sprite.Draw(t, matRight)
}

func (f *Fighter) ResetPosition() {
	f.Position = pixel.V(0, 0)
	f.d = pixel.V(0, 0)
}

func (f *Fighter) Update(leftPressed, rightPressed, upPressed bool) {
	// update the rotation
	if leftPressed {
		f.Rotation += 3
	}
	if rightPressed {
		f.Rotation -= 3
	}

	f.Rotation = f.Rotation % 360
	f.rotationRadians = degreesToRadians(f.Rotation)

	// gravity
	if f.d.Y > maxGravity {
		f.d = f.d.Add(pixel.V(0, gravity))
	}

	// calculate new velocity based on user input
	if upPressed {
		f.d = f.d.Add(pixel.V(-0.35*math.Sin(f.rotationRadians), 0.5*math.Cos(f.rotationRadians)))
	}

	// calculate new position based on velocity
	f.Position = f.Position.Add(f.d)
	f.positionLeft = f.Position.Add(pixel.V(-windowWidth, 0))
	f.positionRight = f.Position.Add(pixel.V(windowWidth, 0))

	// wrap logic
	if f.Position.X > windowWidth {
		f.Position.X = f.Position.X - windowWidth
	}
	if f.Position.X < 0 {
		f.Position.X = f.Position.X + windowWidth
	}
}
