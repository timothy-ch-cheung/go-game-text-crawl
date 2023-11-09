package main

import (
	"fmt"
	img "image"
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/timothy-ch-cheung/go-game-text-crawl/assets"
)

const BTN_SIZE = 20

type GameUI struct {
	ui *ebitenui.UI
}

func NewBtn(icon resource.ImageID, loader *resource.Loader, handler *widget.ButtonClickedHandlerFunc, opts ...widget.WidgetOpt) *widget.Button {
	idle := loadNineSlice(loader.LoadImage(assets.ImgBtnIdle).Data, BTN_SIZE, BTN_SIZE)
	hover := loadNineSlice(loader.LoadImage(assets.ImgBtnHover).Data, BTN_SIZE, BTN_SIZE)
	pressed := loadNineSlice(loader.LoadImage(assets.ImgBtnPressed).Data, BTN_SIZE, BTN_SIZE)
	btnIcon := loader.LoadImage(icon).Data

	button := widget.NewButton(
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    idle,
			Hover:   hover,
			Pressed: pressed,
		}),
		widget.ButtonOpts.Graphic(btnIcon),
		widget.ButtonOpts.ClickedHandler(*handler),
		widget.ButtonOpts.WidgetOpts(opts...),
	)
	button.GraphicImage = &widget.ButtonImageImage{Idle: btnIcon}
	return button
}

func newUI(loader *resource.Loader) *GameUI {
	ui := &ebitenui.UI{}

	rootContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Padding(widget.Insets{
				Bottom: 5,
				Left:   5,
				Right:  5,
				Top:    5,
			}),
			widget.RowLayoutOpts.Spacing(180),
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
		),
		),
	)

	settingContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(widget.RowLayoutOpts.Spacing(442))),
	)

	window := newSettingWindow(loader)

	dialog := NewDialog(
		DialogOpts.DialogImage(&ImageNineSlice{img: loader.LoadImage(assets.ImgFrame).Data, centerWidth: 16, centerHeight: 16}),
		DialogOpts.TextFrameImage(&ImageNineSlice{img: loader.LoadImage(assets.ImgTextFrame).Data, centerWidth: 14, centerHeight: 14}),
		DialogOpts.PlayerPortrait(loader.LoadImage(assets.ImgPortrait).Data),
		DialogOpts.PlayerName("Luna"),
		DialogOpts.FontColor(color.White),
		DialogOpts.TitleFont(loader.LoadFont(assets.FontDefault).Face),
		DialogOpts.TextFont(loader.LoadFont(assets.FontDefault).Face),
		DialogOpts.TextBoxWith(160),
		DialogOpts.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat"),
	)

	var menuHandler widget.ButtonClickedHandlerFunc = func(args *widget.ButtonClickedEventArgs) {
		r := img.Rect(0, 0, 400, 220)
		r = r.Add(img.Point{56, 34})
		window.SetLocation(r)
		ui.AddWindow(window)
	}
	settingContainer.AddChild(NewBtn(assets.ImgIconMenu, loader, &menuHandler))

	var restartHandler widget.ButtonClickedHandlerFunc = func(args *widget.ButtonClickedEventArgs) {
		dialog.RestartDialog()
	}
	settingContainer.AddChild(NewBtn(assets.ImgIconRestart, loader, &restartHandler))

	rootContainer.AddChild(settingContainer)
	rootContainer.AddChild(dialog)

	ui.Container = rootContainer
	return &GameUI{ui: ui}
}

func newSettingWindow(loader *resource.Loader) *widget.Window {
	windowContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(loadNineSlice(loader.LoadImage(assets.ImgFrame).Data, 16, 16)),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(10)),
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(5),
		)),
	)
	window := widget.NewWindow(
		widget.WindowOpts.Contents(windowContainer),
		widget.WindowOpts.Modal(),
		widget.WindowOpts.CloseMode(widget.NONE),
		widget.WindowOpts.MinSize(400, 220),
		widget.WindowOpts.MaxSize(400, 220),
	)

	textInput := widget.NewTextInput(
		widget.TextInputOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  true,
			}),
		),
		widget.TextInputOpts.Padding(widget.NewInsetsSimple(5)),
		widget.TextInputOpts.Image(&widget.TextInputImage{Idle: loadNineSlice(loader.LoadImage(assets.ImgInput).Data, 16, 16)}),
		widget.TextInputOpts.Face(loader.LoadFont(assets.FontDefault).Face),
		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:     color.NRGBA{254, 255, 255, 255},
			Caret:    color.NRGBA{254, 255, 255, 255},
			Disabled: color.NRGBA{125, 125, 125, 255},
		}),
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(loader.LoadFont(assets.FontDefault).Face, 2),
		),
		widget.TextInputOpts.Placeholder("Click here to update dialog"),
	)
	windowContainer.AddChild(textInput)

	textarea := widget.NewTextArea(
		widget.TextAreaOpts.ContainerOpts(
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position:  widget.RowLayoutPositionCenter,
					MaxHeight: 150,
				}),
				widget.WidgetOpts.MinSize(380, 140),
			),
		),
		widget.TextAreaOpts.FontColor(color.Black),
		widget.TextAreaOpts.FontFace(loader.LoadFont(assets.FontDefault).Face),
		widget.TextAreaOpts.ScrollContainerOpts(
			widget.ScrollContainerOpts.Image(&widget.ScrollContainerImage{
				Idle: loadNineSlice(loader.LoadImage(assets.ImgInput).Data, 16, 16),
				Mask: loadNineSlice(loader.LoadImage(assets.ImgInput).Data, 16, 16),
			}),
		),
		widget.TextAreaOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.TextAreaOpts.ControlWidgetSpacing(2),
	)

	textInput.ChangedEvent.AddHandler(func(args interface{}) {
		textarea.SetText(textInput.GetText())
	})

	windowContainer.AddChild(textarea)

	sliderContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(10))),
		widget.ContainerOpts.AutoDisableChildren(),
	)

	sliderLabel := widget.NewLabel(
		widget.LabelOpts.TextOpts(widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		}))),
		widget.LabelOpts.Text("Text Speed:", loader.LoadFont(assets.FontDefault).Face, &widget.LabelColor{Idle: color.Black}),
	)
	sliderContainer.AddChild(sliderLabel)

	var sliderValue *widget.Label

	sliderTrack := loadNineSlice(loader.LoadImage(assets.ImgSliderTrack).Data, 14, 4)
	sliderBtn := loadNineSlice(loader.LoadImage(assets.ImgSliderBtn).Data, 1, 2)
	slider := widget.NewSlider(
		widget.SliderOpts.Direction(widget.DirectionHorizontal),
		widget.SliderOpts.MinMax(1, 10),
		widget.SliderOpts.Images(
			&widget.SliderTrackImage{
				Idle: sliderTrack,
			},
			&widget.ButtonImage{
				Idle:    sliderBtn,
				Pressed: sliderBtn,
			},
		),
		widget.SliderOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(200, 6),
		),
		widget.SliderOpts.FixedHandleSize(4),
		widget.SliderOpts.MinHandleSize(8),
		widget.SliderOpts.TrackOffset(0),
		widget.SliderOpts.PageSizeFunc(func() int {
			return 1
		}),
		widget.SliderOpts.ChangedHandler(func(args *widget.SliderChangedEventArgs) {
			sliderValue.Label = fmt.Sprintf("%d", args.Current)
		}),
	)
	slider.Current = 5
	sliderContainer.AddChild(slider)

	sliderValue = widget.NewLabel(
		widget.LabelOpts.TextOpts(widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Position: widget.RowLayoutPositionCenter,
		}))),
		widget.LabelOpts.Text(fmt.Sprintf("%d", slider.Current), loader.LoadFont(assets.FontDefault).Face, &widget.LabelColor{Idle: color.Black}),
	)
	sliderContainer.AddChild(sliderValue)

	var submitHandler widget.ButtonClickedHandlerFunc = func(args *widget.ButtonClickedEventArgs) {
		window.Close()
	}

	submitLayout := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	submitLayout.AddChild(
		NewBtn(
			assets.ImgIconSubmit, loader, &submitHandler,
			widget.WidgetOpts.LayoutData(
				widget.AnchorLayoutData{HorizontalPosition: widget.AnchorLayoutPositionEnd},
			)))

	bottomPanel := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Spacing(45),
		)),
	)
	bottomPanel.AddChild(sliderContainer)
	bottomPanel.AddChild(submitLayout)
	windowContainer.AddChild(bottomPanel)

	return window
}

func (game *GameUI) update() {
	game.ui.Update()
}

func (game *GameUI) draw(screen *ebiten.Image) {
	game.ui.Draw(screen)
}
