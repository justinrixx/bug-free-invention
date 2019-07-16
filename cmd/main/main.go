package main

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "smashteroids",
		Bounds: pixel.R(0, 0, 512, 512),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.SetSmooth(true)

	pic, err := loadPicture("hiking.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds() /*pixel.R(0, 0, 16, 16)*/)

	position := win.Bounds().Center()
	d := pixel.V(3, 0)
	// mat = mat.Scaled(win.Bounds().Center(), 0.25)

	// imd := imdraw.New(nil)
	// imd.Precision = 32

	// rect := smashteroids.Rectangle{
	// 	Color:  colornames.Black,
	// 	Rect:   pixel.R(0, 0, 32, 32),
	// 	Matrix: pixel.IM.Moved(win.Bounds().Center()),
	// }

	// last := time.Now()
	for !win.Closed() {
		// dt = time.Since(last).Seconds()
		// last = time.Now()

		// logic
		if position.X+100 >= 512 || position.X-100 <= 0 {
			d.X *= -1
		}

		position = position.Add(d)

		mat := pixel.IM.Moved(position).Scaled(position, .25)

		win.Clear(colornames.Azure)

		// mat = mat.Moved(pixel.V(1, 0))
		sprite.Draw(win, mat)

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
