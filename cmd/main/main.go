package main

import (
	"fmt"
	"image"
	_ "image/png"
	"io/ioutil"
	"math"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

const (
	radianScaleFactor = math.Pi / 180

	windowWidth  = 1024
	windowHeight = 764

	shipWidth     = 50
	halfShipWidth = shipWidth / 2

	shipHeight     = 60
	halfShipHeight = shipHeight / 2
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "smashteroids",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("images/ship.png")
	if err != nil {
		panic(err)
	}

	face, err := loadTTF("fonts/Hyperspace.otf", 32)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)
	txt := text.New(pixel.V(100, 100), atlas)

	fmt.Fprintln(txt, "bacon")

	sprite := pixel.NewSprite(pic, pic.Bounds())

	position := win.Bounds().Center()
	var positionAbove, positionBelow, positionLeft, positionRight pixel.Vec
	d := pixel.V(0, 0)
	rotation := 0

	// last := time.Now()
	for !win.Closed() {
		// dt = time.Since(last).Seconds()
		// last = time.Now()

		// wrap logic
		if position.X > windowWidth {
			position.X = position.X - windowWidth
		}
		if position.X < 0 {
			position.X = position.X + windowWidth
		}
		if position.Y > windowHeight {
			position.Y = position.Y - windowHeight
		}
		if position.Y < 0 {
			position.Y = position.Y + windowHeight
		}

		// calculate new velocity based on user input
		if win.Pressed(pixelgl.KeyLeft) {
			rotation += 3
		}
		if win.Pressed(pixelgl.KeyRight) {
			rotation -= 3
		}

		rotation = rotation % 360
		rotationRadians := degreesToRadians(rotation)

		if win.Pressed(pixelgl.KeyUp) {
			d = d.Add(pixel.V(-0.25*math.Sin(rotationRadians), 0.25*math.Cos(rotationRadians)))
		}

		// calculate new position based on velocity
		position = position.Add(d)
		positionAbove = position.Add(pixel.V(0, windowHeight))
		positionBelow = position.Add(pixel.V(0, -windowHeight))
		positionLeft = position.Add(pixel.V(-windowWidth, 0))
		positionRight = position.Add(pixel.V(windowWidth, 0))

		mat := pixel.IM.Moved(position).Scaled(position, 2).Rotated(position, rotationRadians)
		matAbove := pixel.IM.Moved(positionAbove).Scaled(positionAbove, 2).Rotated(positionAbove, rotationRadians)
		matBelow := pixel.IM.Moved(positionBelow).Scaled(positionBelow, 2).Rotated(positionBelow, rotationRadians)
		matLeft := pixel.IM.Moved(positionLeft).Scaled(positionLeft, 2).Rotated(positionLeft, rotationRadians)
		matRight := pixel.IM.Moved(positionRight).Scaled(positionRight, 2).Rotated(positionRight, rotationRadians)

		win.Clear(colornames.Black)

		sprite.Draw(win, mat)
		sprite.Draw(win, matAbove)
		sprite.Draw(win, matBelow)
		sprite.Draw(win, matLeft)
		sprite.Draw(win, matRight)

		txt.Draw(win, pixel.IM)

		win.Update()
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func degreesToRadians(degrees int) float64 {
	return float64(degrees) * radianScaleFactor
}
