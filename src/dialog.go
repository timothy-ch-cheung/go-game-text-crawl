package main

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type Dialog struct {
	init           *widget.MultiOnce
	container      *widget.Container
	dialogImage    *image.NineSlice
	fontColor      color.Color
	textFont       font.Face
	playerNameFont font.Face
	playerName     string
	playerPortrait *ebiten.Image
	textBoxWidth   int
	textBoxHeight  int
}

type DialogOpt func(dialog *Dialog)

type DialogOptions struct {
}

var DialogOpts DialogOptions

func NewDialog(opts ...DialogOpt) *Dialog {
	dialog := &Dialog{
		init: &widget.MultiOnce{},
	}

	dialog.init.Append(dialog.createWidget)

	for _, opt := range opts {
		opt(dialog)
	}

	return dialog
}

func (dialog *Dialog) GetWidget() *widget.Widget {
	dialog.init.Do()
	return dialog.container.GetWidget()
}

func (dialog *Dialog) PreferredSize() (int, int) {
	dialog.init.Do()
	return dialog.container.PreferredSize()
}

func (dialog *Dialog) createWidget() {
	dialog.container = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(5),
		)),
		widget.ContainerOpts.BackgroundImage(dialog.dialogImage),
	)

	graphic := widget.NewGraphic(widget.GraphicOpts.Image(dialog.playerPortrait))
	dialog.container.AddChild(graphic)

	textContiner := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(5),
		)),
	)
	dialog.container.AddChild(textContiner)

	label := widget.NewLabel(
		widget.LabelOpts.Text(dialog.playerName, dialog.playerNameFont, &widget.LabelColor{Idle: dialog.fontColor}),
	)
	textContiner.AddChild(label)

	textArea := widget.NewTextArea(
		widget.TextAreaOpts.ContainerOpts(
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.MinSize(dialog.textBoxWidth, dialog.textBoxHeight),
			),
		),
		widget.TextAreaOpts.Text(""),
		widget.TextAreaOpts.FontFace(dialog.textFont),
		widget.TextAreaOpts.FontColor(dialog.fontColor),
	)
	textContiner.AddChild(textArea)
}

func (option DialogOptions) PlayerPortrait(image *ebiten.Image) DialogOpt {
	return func(dialog *Dialog) {
		dialog.playerPortrait = image
	}
}

func (option DialogOptions) PlayerName(name string) DialogOpt {
	return func(dialog *Dialog) {
		dialog.playerName = name
	}
}

func (option DialogOptions) FontColor(color color.Color) DialogOpt {
	return func(dialog *Dialog) {
		dialog.fontColor = color
	}
}

func (option DialogOptions) TitleFont(titleFont font.Face) DialogOpt {
	return func(dialog *Dialog) {
		dialog.playerNameFont = titleFont
	}
}

func (option DialogOptions) TextFont(textFont font.Face) DialogOpt {
	return func(dialog *Dialog) {
		dialog.textFont = textFont
	}
}

func (option DialogOptions) DialogImage(dialogImage *image.NineSlice) DialogOpt {
	return func(dialog *Dialog) {
		dialog.dialogImage = dialogImage
	}
}

func (option DialogOptions) TextBoxSize(width int, height int) DialogOpt {
	return func(dialog *Dialog) {
		dialog.textBoxWidth = width
		dialog.textBoxHeight = height
	}
}
