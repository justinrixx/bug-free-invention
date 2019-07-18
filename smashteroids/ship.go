package smashteroids

import (
	"math"

	"github.com/faiface/pixel"
)

type Ship struct {
	Position pixel.Vec
	Rotation int
	Sprite   *pixel.Sprite

	d pixel.Vec

	positionLeft  pixel.Vec
	positionRight pixel.Vec

	rotationRadians float64
}

func (s *Ship) Draw(t pixel.Target) {
	mat := pixel.IM.Moved(s.Position).Rotated(s.Position, s.rotationRadians)
	matLeft := pixel.IM.Moved(s.positionLeft).Rotated(s.positionLeft, s.rotationRadians)
	matRight := pixel.IM.Moved(s.positionRight).Rotated(s.positionRight, s.rotationRadians)

	s.Sprite.Draw(t, mat)
	s.Sprite.Draw(t, matLeft)
	s.Sprite.Draw(t, matRight)
}

func (s *Ship) Update(leftPressed, rightPressed, upPressed bool) {
	// update the rotation
	if leftPressed {
		s.Rotation += 3
	}
	if rightPressed {
		s.Rotation -= 3
	}

	s.Rotation = s.Rotation % 360
	s.rotationRadians = degreesToRadians(s.Rotation)

	// gravity
	if s.d.Y > maxGravity {
		s.d = s.d.Add(pixel.V(0, gravity))
	}

	// calculate new velocity based on user input
	if upPressed {
		s.d = s.d.Add(pixel.V(-0.35*math.Sin(s.rotationRadians), 0.5*math.Cos(s.rotationRadians)))
	}

	// calculate new position based on velocity
	s.Position = s.Position.Add(s.d)
	s.positionLeft = s.Position.Add(pixel.V(-windowWidth, 0))
	s.positionRight = s.Position.Add(pixel.V(windowWidth, 0))

	// wrap logic
	if s.Position.X > windowWidth {
		s.Position.X = s.Position.X - windowWidth
	}
	if s.Position.X < 0 {
		s.Position.X = s.Position.X + windowWidth
	}
}
