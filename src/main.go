package main

import (
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/timothy-ch-cheung/go-game-text-crawl/assets"
)

const (
	SCREEN_WIDTH  = 512
	SCREEN_HEIGHT = 288
	SCALE         = 2
)

type Game struct {
	loader *resource.Loader
	ui     *ebitenui.UI
}

func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.loader.LoadImage(assets.ImgBackground).Data, &ebiten.DrawImageOptions{})
	g.ui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func newGame() *Game {
	audioContext := audio.NewContext(44100)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAssetFunc
	assets.RegisterImageResources(loader)

	return &Game{
		loader: loader,
		ui:     newUI(loader),
	}
}

func main() {
	ebiten.SetWindowSize(SCREEN_WIDTH*SCALE, SCREEN_HEIGHT*SCALE)
	ebiten.SetWindowTitle("Text Crawl")

	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
