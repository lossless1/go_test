package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type ImageI interface {
	ColorModel() color.Model
	// Bounds returns the domain for which At can return non-zero color.
	// The bounds do not necessarily contain the point (0, 0).
	Bounds() image.Rectangle
	// At returns the color of the pixel at (x, y).
	// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
	// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
	At(x, y int) color.Color
}

type Image struct {
	width  int
	height int
}

// func (img Image) ShowImage(x, y int) (int, int) {
// 	return 1, 2
// }

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	img_func := func(x, y int) uint8 {
		//return uint8(x*y)
		//return uint8((x+y) / 2)
		return uint8(x ^ y)
	}
	v := img_func(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 64}
	fmt.Print(m)
	pic.ShowImage(m)
}
