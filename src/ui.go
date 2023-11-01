package main

import (
	img "image"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/timothy-ch-cheung/go-game-text-crawl/assets"
)

const BTN_SIZE = 20

func loadImageNineSlice(img *ebiten.Image, centerWidth int, centerHeight int) *image.NineSlice {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	return image.NewNineSlice(img,
		[3]int{(width - centerWidth) / 2, centerWidth, width - (width-centerWidth)/2 - centerWidth},
		[3]int{(height - centerHeight) / 2, centerHeight, height - (height-centerHeight)/2 - centerHeight},
	)
}

func NewBtn(icon resource.ImageID, loader *resource.Loader, handler *widget.ButtonClickedHandlerFunc) *widget.Button {
	idle := loadImageNineSlice(loader.LoadImage(assets.ImgBtnIdle).Data, BTN_SIZE, BTN_SIZE)
	hover := loadImageNineSlice(loader.LoadImage(assets.ImgBtnHover).Data, BTN_SIZE, BTN_SIZE)
	pressed := loadImageNineSlice(loader.LoadImage(assets.ImgBtnPressed).Data, BTN_SIZE, BTN_SIZE)
	btnIcon := loader.LoadImage(icon).Data

	button := widget.NewButton(
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    idle,
			Hover:   hover,
			Pressed: pressed,
		}),
		widget.ButtonOpts.Graphic(btnIcon),
		widget.ButtonOpts.ClickedHandler(*handler),
	)
	button.GraphicImage = &widget.ButtonImageImage{Idle: btnIcon}
	return button
}

func newUI(loader *resource.Loader) *ebitenui.UI {
	ui := &ebitenui.UI{}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(widget.RowLayoutOpts.Padding(widget.Insets{
			Bottom: 5,
			Left:   5,
			Right:  5,
			Top:    5,
		}), widget.RowLayoutOpts.Spacing(442))),
	)

	windowContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(loadImageNineSlice(loader.LoadImage(assets.ImgFrame).Data, 16, 16)),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	window := widget.NewWindow(
		widget.WindowOpts.Contents(windowContainer),
		widget.WindowOpts.Modal(),
		widget.WindowOpts.CloseMode(widget.CLICK_OUT),
		widget.WindowOpts.MinSize(400, 220),
		widget.WindowOpts.MaxSize(400, 220),
	)

	var menuHandler widget.ButtonClickedHandlerFunc = func(args *widget.ButtonClickedEventArgs) {
		r := img.Rect(0, 0, 400, 220)
		r = r.Add(img.Point{56, 34})
		window.SetLocation(r)
		ui.AddWindow(window)
	}
	rootContainer.AddChild(NewBtn(assets.ImgIconMenu, loader, &menuHandler))

	var restartHandler widget.ButtonClickedHandlerFunc = func(args *widget.ButtonClickedEventArgs) {

	}
	rootContainer.AddChild(NewBtn(assets.ImgIconRestart, loader, &restartHandler))

	ui.Container = rootContainer
	return ui
}
