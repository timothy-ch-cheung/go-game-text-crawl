package main

import (
	img "image"
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

type Dialog struct {
	init           *widget.MultiOnce
	container      *widget.Container
	dialogImage    *ImageNineSlice
	fontColor      color.Color
	textFont       font.Face
	playerNameFont font.Face
	playerName     string
	playerPortrait *ebiten.Image
	textBoxWidth   int
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

func (dialog *Dialog) SetLocation(rect img.Rectangle) {
	dialog.init.Do()
	dialog.container.GetWidget().Rect = rect
}

func (dialog *Dialog) Render(screen *ebiten.Image, def widget.DeferredRenderFunc) {
	dialog.init.Do()
	dialog.container.Render(screen, def)
}

func (dialog *Dialog) createWidget() {
	dialogImage := loadImageNineSlice(*dialog.dialogImage)
	portraitWidth := dialog.dialogImage.img.Bounds().Dx()
	portraitHeight := dialog.dialogImage.img.Bounds().Dy()
	xPadding := (portraitWidth - dialog.dialogImage.centerWidth) / 2
	yPadding := (portraitHeight - dialog.dialogImage.centerHeight) / 2

	dialog.container = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Left:   xPadding,
				Right:  xPadding,
				Top:    yPadding,
				Bottom: yPadding,
			}),
		)),
		widget.ContainerOpts.BackgroundImage(dialogImage),
	)

	graphic := widget.NewGraphic(widget.GraphicOpts.Image(dialog.playerPortrait))
	dialog.container.AddChild(graphic)

	textContiner := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{Left: 2, Right: 2, Top: 2, Bottom: 2}),
		)),
	)
	dialog.container.AddChild(textContiner)

	label := widget.NewLabel(
		widget.LabelOpts.Text(dialog.playerName, dialog.playerNameFont, &widget.LabelColor{Idle: dialog.fontColor}),
	)
	textContiner.AddChild(label)

	textBoxHeight := portraitHeight - label.GetWidget().MinHeight
	textArea := widget.NewTextArea(
		widget.TextAreaOpts.ContainerOpts(
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.MinSize(dialog.textBoxWidth, textBoxHeight),
			),
		),
		widget.TextAreaOpts.ScrollContainerOpts(
			widget.ScrollContainerOpts.Image(&widget.ScrollContainerImage{
				Idle: image.NewNineSliceColor(color.Transparent),
				Mask: image.NewNineSliceColor(color.Transparent),
			}),
		),
		widget.TextAreaOpts.Text("Lorem ipsum dolor sit amet"),
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

func (option DialogOptions) DialogImage(dialogImage *ImageNineSlice) DialogOpt {
	return func(dialog *Dialog) {
		dialog.dialogImage = dialogImage
	}
}

func (option DialogOptions) TextBoxWith(width int) DialogOpt {
	return func(dialog *Dialog) {
		dialog.textBoxWidth = width
	}
}
