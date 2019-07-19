package main

import (
	"image"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/golang/freetype/truetype"
	"github.com/justinrixx/bug-free-invention/smashteroids"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
)

const (
	windowWidth  = 1024
	windowHeight = 764

	player1 = 1
)

func main() {
	pixelgl.Run(run)
}

func run() {
	// make the window
	cfg := pixelgl.WindowConfig{
		Title:  "SMASHTEROIDS",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// load assets
	pic, err := loadPicture("images/ship.png")
	if err != nil {
		panic(err)
	}
	face, err := loadTTF("fonts/Hyperspace.otf", 24)
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	position := win.Bounds().Center()
	rotation := 0

	player := smashteroids.Player{
		Ship: &smashteroids.Fighter{
			Position: position,
			Rotation: rotation,
			Sprite:   sprite,
		},
		Lives:         5,
		Fuel:          500.0,
		TeamID:        player1,
		FuelLocation:  pixel.V(25, windowHeight-36),
		LivesLocation: pixel.V(25, windowHeight-65),
	}

	player.Initialize(face)

	terrain := smashteroids.GenerateTerrain(128, 100, 75, .7)

	// last := time.Now()
	for !win.Closed() {
		// dt = time.Since(last).Seconds()
		// last = time.Now()

		player.Update(win.Pressed(pixelgl.KeyLeft), win.Pressed(pixelgl.KeyRight), win.Pressed(pixelgl.KeyUp))

		win.Clear(colornames.Black)

		player.Draw(win)

		smashteroids.DrawTerrain(terrain, win, windowWidth)

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
