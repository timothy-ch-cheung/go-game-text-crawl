package main

import (
	img "image"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type HideableGraphic struct {
	init          *widget.MultiOnce
	container     *widget.Container
	containerOpts []widget.ContainerOpt
	padding       widget.Insets
	graphic       *ebiten.Image
	hidden        bool
}

type HideableGraphicOpt func(hideableGraphic *HideableGraphic)

type HideableGraphicOptions struct {
}

var HideableGraphicOpts HideableGraphicOptions

/////////////////////
// Widget Creation //
/////////////////////

func NewHideableGraphic(opts ...HideableGraphicOpt) *HideableGraphic {
	hideableGraphic := &HideableGraphic{
		init:   &widget.MultiOnce{},
		hidden: false,
	}

	hideableGraphic.init.Append(hideableGraphic.createWidget)

	for _, opt := range opts {
		opt(hideableGraphic)
	}

	return hideableGraphic
}

func (hideableGraphic *HideableGraphic) createWidget() {
	hideableGraphic.container = widget.NewContainer(
		append(hideableGraphic.containerOpts, widget.ContainerOpts.Layout(
			widget.NewAnchorLayout(widget.AnchorLayoutOpts.Padding(hideableGraphic.padding)),
		))...,
	)

	graphic := widget.NewGraphic(widget.GraphicOpts.Image(hideableGraphic.graphic))

	hideableGraphic.container.AddChild(graphic)
}

////////////////////////////
// Initial Widget Options //
////////////////////////////

func (option HideableGraphicOptions) Graphic(image *ebiten.Image) HideableGraphicOpt {
	return func(hideableGraphic *HideableGraphic) {
		hideableGraphic.graphic = image
	}
}

func (option HideableGraphicOptions) Padding(padding widget.Insets) HideableGraphicOpt {
	return func(hideableGraphic *HideableGraphic) {
		hideableGraphic.padding = padding
	}
}

func (options HideableGraphicOptions) ContainerOpts(opts ...widget.ContainerOpt) HideableGraphicOpt {
	return func(hideableGraphic *HideableGraphic) {
		hideableGraphic.containerOpts = append(hideableGraphic.containerOpts, opts...)
	}
}

///////////////////////////
// Active Widget Options //
///////////////////////////

func (hideableGraphic *HideableGraphic) Hide() {
	hideableGraphic.hidden = true
}

func (hideableGraphic *HideableGraphic) Show() {
	hideableGraphic.hidden = false
}

func (hideableGraphic *HideableGraphic) IsHidden() bool {
	return hideableGraphic.hidden
}

/////////////////////////////////////////////
// Implement PreferredSizeLocateableWidget //
/////////////////////////////////////////////

func (hideableGraphic *HideableGraphic) GetWidget() *widget.Widget {
	hideableGraphic.init.Do()
	return hideableGraphic.container.GetWidget()
}

func (hideableGraphic *HideableGraphic) PreferredSize() (int, int) {
	hideableGraphic.init.Do()
	return hideableGraphic.container.PreferredSize()
}

func (hideableGraphic *HideableGraphic) SetLocation(rect img.Rectangle) {
	hideableGraphic.init.Do()
	hideableGraphic.container.GetWidget().Rect = rect
}

func (hideableGraphic *HideableGraphic) Render(screen *ebiten.Image, def widget.DeferredRenderFunc) {
	hideableGraphic.init.Do()
	if !hideableGraphic.hidden {
		hideableGraphic.container.Render(screen, def)
	}
}
