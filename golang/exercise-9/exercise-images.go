package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	rect := image.Rect(0, 0, 100, 100)
	return rect
}

func (i Image) At(x int, y int) color.Color {
	v := uint8((x - y) ^ (x + y))
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
