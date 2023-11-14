package main

import (
	"fmt"
	img "image"
	"image/color"
	"time"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

const (
	BASE_SPEED               = 2
	DEFAULT_SPEED_MULTIPLIER = 5
	PAGE_END_BLINK_SPEED     = 3
)

var lastUpdated time.Time
var lastBlinked time.Time

type Dialog struct {
	init                    *widget.MultiOnce
	container               *widget.Container
	dialogImage             *ImageNineSlice
	fontColor               color.Color
	textFont                font.Face
	playerNameFont          font.Face
	playerName              string
	playerPortrait          *ebiten.Image
	textBoxWidth            int
	textBoxHeight           int
	text                    string
	dialogPage              *DialogPage
	speedMultiplier         float64
	setText                 func(string)
	pageEndIndicator        *ebiten.Image
	pageEndIndicatorGraphic *HideableGraphic
	completed               bool
}

type DialogOpt func(dialog *Dialog)

type DialogOptions struct {
}

type DialogPage struct {
	textGroups       []string
	currentPage      int
	currentCharacter int
}

var DialogOpts DialogOptions

/////////////////////
// Widget Creation //
/////////////////////

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

func (dialog *Dialog) createWidget() {
	dialog.speedMultiplier = DEFAULT_SPEED_MULTIPLIER
	lastUpdated = time.Now()
	dialogImage := loadImageNineSlice(*dialog.dialogImage)
	xPadding, yPadding := getPadding(*dialog.dialogImage)

	dialog.container = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewStackedLayout()),
	)

	mainDialog := widget.NewContainer(
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
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(dialog.textBoxWidth+dialog.playerPortrait.Bounds().Dx(), 50),
		),
	)
	dialog.container.AddChild(mainDialog)

	graphic := widget.NewGraphic(widget.GraphicOpts.Image(dialog.playerPortrait))
	mainDialog.AddChild(graphic)

	textContiner := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{Left: 2, Right: 2, Top: 2, Bottom: 2}),
			widget.RowLayoutOpts.Spacing(2),
		)),
	)
	mainDialog.AddChild(textContiner)

	label := widget.NewLabel(
		widget.LabelOpts.Text(dialog.playerName, dialog.playerNameFont, &widget.LabelColor{Idle: dialog.fontColor}),
	)
	textContiner.AddChild(label)

	_, labelHeight := label.PreferredSize()
	dialog.textBoxHeight = dialog.playerPortrait.Bounds().Dy() - labelHeight - 6
	dialog.dialogPage = &DialogPage{
		textGroups:       GroupText(dialog.textFont, dialog.text, dialog.textBoxWidth, dialog.textBoxHeight),
		currentPage:      0,
		currentCharacter: 0,
	}
	textArea := widget.NewTextArea(
		widget.TextAreaOpts.ContainerOpts(
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.RowLayoutData{
					Position:  widget.RowLayoutPositionCenter,
					MaxWidth:  dialog.textBoxWidth,
					MaxHeight: dialog.textBoxHeight,
				}),
				widget.WidgetOpts.MinSize(dialog.textBoxWidth, dialog.textBoxHeight),
			),
		),
		widget.TextAreaOpts.ScrollContainerOpts(
			widget.ScrollContainerOpts.Image(&widget.ScrollContainerImage{
				Idle: image.NewNineSliceColor(color.Transparent),
				Mask: image.NewNineSliceColor(color.Opaque),
			}),
		),
		widget.TextAreaOpts.Text(dialog.dialogPage.GetCurrentText()),
		widget.TextAreaOpts.FontFace(dialog.textFont),
		widget.TextAreaOpts.FontColor(dialog.fontColor),
		widget.TextAreaOpts.TextPadding(widget.NewInsetsSimple(2)),
	)
	textContiner.AddChild(textArea)
	dialog.setText = func(text string) {
		textArea.SetText(text)
	}

	pageEndContainer := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewAnchorLayout()))
	dialog.container.AddChild(pageEndContainer)

	dialog.pageEndIndicatorGraphic = NewHideableGraphic(
		HideableGraphicOpts.Graphic(dialog.pageEndIndicator),
		HideableGraphicOpts.Padding(widget.NewInsetsSimple(2)),
		HideableGraphicOpts.ContainerOpts(
			widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionEnd,
			}),
			),
		))
	dialog.pageEndIndicatorGraphic.Hide()
	pageEndContainer.AddChild(dialog.pageEndIndicatorGraphic)
}

////////////////////////////
// Initial Widget Options //
////////////////////////////

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

func (option DialogOptions) Text(text string) DialogOpt {
	return func(dialog *Dialog) {
		dialog.text = text
	}
}

func (option DialogOptions) PageEndIndicator(image *ebiten.Image) DialogOpt {
	return func(dialog *Dialog) {
		dialog.pageEndIndicator = image
	}
}

///////////////////////////
// Active Widget Options //
///////////////////////////

func (dialog *Dialog) SetSpeedMultiplier(multiplier float64) {
	dialog.speedMultiplier = multiplier
}

func (dialog *Dialog) AdvanceDialog() {
	if !dialog.dialogPage.IsPageEnd() {
		dialog.dialogPage.currentCharacter = len(dialog.dialogPage.GetCurrentText()) - 2
	} else if !dialog.dialogPage.IsLastPage() {
		dialog.dialogPage.currentCharacter = 0
		dialog.dialogPage.currentPage += 1
		lastUpdated = time.Now()
	} else {
		dialog.setText("")
		dialog.completed = true
	}
	dialog.pageEndIndicatorGraphic.Hide()
}

func (dialog *Dialog) RestartDialog() {
	dialog.dialogPage.currentPage = 0
	dialog.dialogPage.currentCharacter = 0
	dialog.completed = false
	dialog.pageEndIndicatorGraphic.Hide()
	lastUpdated = time.Now()
}

func (dialog *Dialog) SetText(text string) {
	dialog.dialogPage.currentPage = 0
	dialog.dialogPage.currentCharacter = 0
	dialog.completed = false
	dialog.pageEndIndicatorGraphic.Hide()
	dialog.dialogPage.textGroups = GroupText(dialog.textFont, text, dialog.textBoxWidth, dialog.textBoxHeight)
	lastUpdated = time.Now()
}

/////////////////////////////////////////////
// Implement PreferredSizeLocateableWidget //
/////////////////////////////////////////////

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
	now := time.Now()
	if !dialog.dialogPage.IsPageEnd() && now.Sub(lastUpdated).Seconds() > 1/(dialog.speedMultiplier*BASE_SPEED) {
		dialog.dialogPage.currentCharacter += 1
		dialog.setText(dialog.dialogPage.GetDisplayedText())
		lastUpdated = now
	}
	if dialog.dialogPage.IsPageEnd() && !dialog.completed && now.Sub(lastBlinked).Seconds() > 1/(float64(PAGE_END_BLINK_SPEED)) {
		if dialog.pageEndIndicatorGraphic.IsHidden() {
			dialog.pageEndIndicatorGraphic.Show()
		} else {
			dialog.pageEndIndicatorGraphic.Hide()
		}
		lastBlinked = now
	}

	dialog.init.Do()
	dialog.container.Render(screen, def)
}

//////////////////////
// Helper Functions //
//////////////////////

func getPadding(imgNineSlice ImageNineSlice) (int, int) {
	portraitWidth := imgNineSlice.img.Bounds().Dx()
	portraitHeight := imgNineSlice.img.Bounds().Dy()
	xPadding := (portraitWidth - imgNineSlice.centerWidth) / 2
	yPadding := (portraitHeight - imgNineSlice.centerHeight) / 2
	return xPadding, yPadding
}

func replaceAfterPosition(input string, position int) string {
	runes := []rune(input)
	for i := position + 1; i < len(runes); i++ {
		if runes[i] != ' ' {
			runes[i] = '\u00A0'
		}
	}
	return string(runes)
}

func (page DialogPage) String() string {
	return fmt.Sprintf("Current Page: [%d], Current Character: [%d],\nText Groups: %v", page.currentPage, page.currentCharacter, page.textGroups)
}

func (dialogPage *DialogPage) GetCurrentText() string {
	return dialogPage.textGroups[dialogPage.currentPage]
}

func (dialogPage *DialogPage) GetDisplayedText() string {
	currentPage := dialogPage.textGroups[dialogPage.currentPage]
	return replaceAfterPosition(currentPage, dialogPage.currentCharacter)
}

func (dialogPage *DialogPage) IsPageEnd() bool {
	return dialogPage.currentCharacter == len(dialogPage.textGroups[dialogPage.currentPage])-1
}

func (dialogPage *DialogPage) IsLastPage() bool {
	return dialogPage.currentPage == len(dialogPage.textGroups)-1
}
