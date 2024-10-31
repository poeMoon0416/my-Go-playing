package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type MyImage struct {
	w, h int
	// blue scale [][]uint8
	v [][]uint8
}

func (mi MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, mi.w, mi.h)
}

func (mi MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (mi MyImage) At(x, y int) color.Color {
	v := mi.v[y][x]
	return color.RGBA{uint8(v), uint8(v), 255, 255}
}

func Pic(dx, dy int) image.Image {
	// make v
	v := make([][]uint8, dy)
	for i := range v {
		v[i] = make([]uint8, dx)
	}
	for i := range v {
		for j := range v[i] {
			v[i][j] = uint8((i + j) / 2)
		}
	}
	// make MyImage
	mi := MyImage{w: dx, h: dy, v: v}
	return mi
}

func main() {
	pic.ShowImage(Pic(512, 512))
}
