package main

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
)

func newUI() *ebitenui.UI {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(widget.AnchorLayoutOpts.Padding(widget.Insets{
			Bottom: 5,
			Left:   5,
			Right:  5,
		}))),
	)

	return &ebitenui.UI{
		Container: rootContainer,
	}
}
