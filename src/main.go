package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/timothy-ch-cheung/go-game-text-crawl/assets"
)

const (
	SCREEN_WIDTH  = 512
	SCREEN_HEIGHT = 288
	SCALE         = 2
)

type Game struct {
	loader       *resource.Loader
	ui           *GameUI
	inputSystem  input.System
	inputHandler input.Handler
}

const (
	ActionAdvanceDialog input.Action = iota
)

func (g *Game) Update() error {
	g.inputSystem.Update()
	g.ui.update()

	if g.inputHandler.ActionIsJustPressed(ActionAdvanceDialog) {
		g.ui.AdvanceDialog()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.loader.LoadImage(assets.ImgBackground).Data, &ebiten.DrawImageOptions{})
	g.ui.draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func newGame() *Game {
	game := &Game{}

	game.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})

	keymap := input.Keymap{ActionAdvanceDialog: {input.KeyA, input.KeySpace}}
	game.inputHandler = *game.inputSystem.NewHandler(0, keymap)

	audioContext := audio.NewContext(44100)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAssetFunc
	game.loader = loader

	assets.RegisterImageResources(loader)
	assets.RegisterFontResources(loader)

	game.ui = newUI(loader)

	return game
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH*SCALE, SCREEN_HEIGHT*SCALE)
	ebiten.SetWindowTitle("Text Crawl")

	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
