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

	positionAbove pixel.Vec
	positionBelow pixel.Vec
	positionLeft  pixel.Vec
	positionRight pixel.Vec

	rotationRadians float64
}

func (s *Ship) Draw(t pixel.Target) {
	mat := pixel.IM.Moved(s.Position).Scaled(s.Position, 2).Rotated(s.Position, s.rotationRadians)
	matAbove := pixel.IM.Moved(s.positionAbove).Scaled(s.positionAbove, 2).Rotated(s.positionAbove, s.rotationRadians)
	matBelow := pixel.IM.Moved(s.positionBelow).Scaled(s.positionBelow, 2).Rotated(s.positionBelow, s.rotationRadians)
	matLeft := pixel.IM.Moved(s.positionLeft).Scaled(s.positionLeft, 2).Rotated(s.positionLeft, s.rotationRadians)
	matRight := pixel.IM.Moved(s.positionRight).Scaled(s.positionRight, 2).Rotated(s.positionRight, s.rotationRadians)

	s.Sprite.Draw(t, mat)
	s.Sprite.Draw(t, matAbove)
	s.Sprite.Draw(t, matBelow)
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

	// calculate new velocity based on user input
	if upPressed {
		s.d = s.d.Add(pixel.V(-0.25*math.Sin(s.rotationRadians), 0.25*math.Cos(s.rotationRadians)))
	}

	// calculate new position based on velocity
	s.Position = s.Position.Add(s.d)
	s.positionAbove = s.Position.Add(pixel.V(0, windowHeight))
	s.positionBelow = s.Position.Add(pixel.V(0, -windowHeight))
	s.positionLeft = s.Position.Add(pixel.V(-windowWidth, 0))
	s.positionRight = s.Position.Add(pixel.V(windowWidth, 0))

	// wrap logic
	if s.Position.X > windowWidth {
		s.Position.X = s.Position.X - windowWidth
	}
	if s.Position.X < 0 {
		s.Position.X = s.Position.X + windowWidth
	}
	if s.Position.Y > windowHeight {
		s.Position.Y = s.Position.Y - windowHeight
	}
	if s.Position.Y < 0 {
		s.Position.Y = s.Position.Y + windowHeight
	}
}
