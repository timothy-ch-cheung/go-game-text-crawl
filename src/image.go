package main

import (
	"github.com/ebitenui/ebitenui/image"
	"github.com/hajimehoshi/ebiten/v2"
)

type ImageNineSlice struct {
	img          *ebiten.Image
	centerWidth  int
	centerHeight int
}

func loadImageNineSlice(imageNineSlice ImageNineSlice) *image.NineSlice {
	return loadNineSlice(imageNineSlice.img, imageNineSlice.centerWidth, imageNineSlice.centerHeight)
}

func loadNineSlice(img *ebiten.Image, centerWidth int, centerHeight int) *image.NineSlice {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	return image.NewNineSlice(img,
		[3]int{(width - centerWidth) / 2, centerWidth, width - (width-centerWidth)/2 - centerWidth},
		[3]int{(height - centerHeight) / 2, centerHeight, height - (height-centerHeight)/2 - centerHeight},
	)
}
