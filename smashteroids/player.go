package smashteroids

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

var (
	atlas *text.Atlas
)

type Player struct {
	Ship   Ship
	Lives  int
	Fuel   float64
	TeamID int
	// Bullets
	FuelLocation  pixel.Vec
	LivesLocation pixel.Vec
}

func (p *Player) Initialize(face font.Face) {
	atlas = text.NewAtlas(face, text.ASCII)
}

func (p *Player) Update(leftPressed, rightPressed, upPressed bool) {
	hasFuel := false
	if p.Fuel > 0 {
		hasFuel = true
		if upPressed {
			p.Fuel -= p.Ship.GetFuelUse()
		}
	}
	p.Ship.Update(leftPressed, rightPressed, upPressed && hasFuel)
}

func (p *Player) Draw(t pixel.Target) {
	p.Ship.Draw(t)

	fuel := text.New(p.FuelLocation, atlas)
	lives := text.New(p.LivesLocation, atlas)

	fmt.Fprintln(fuel, "fuel")
	fmt.Fprintf(lives, "lives:%d", p.Lives)
	fuel.Draw(t, pixel.IM)
	lives.Draw(t, pixel.IM)

	p.DrawFuel(t)
}

func (p *Player) DrawFuel(t pixel.Target) {
	imd := imdraw.New(nil)

	imd.Color = colornames.White

	// outside frame
	imd.Push(p.FuelLocation.Add(pixel.V(75, 0)))
	imd.Push(p.FuelLocation.Add(pixel.V(150, 18)))
	imd.Rectangle(1)

	// inside
	imd.Push(p.FuelLocation.Add(pixel.V(75, 0)))
	imd.Push(p.FuelLocation.Add(pixel.V(75+(75*p.Fuel/fuelCapacity), 18)))
	imd.Rectangle(0)

	imd.Draw(t)
}

func (p *Player) Kill() {
	p.Lives--

	// move back to the middle
	p.Ship.ResetPosition()
}
