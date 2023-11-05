package assets

import (
	"embed"
	"io"

	_ "image/png"

	resource "github.com/quasilyte/ebitengine-resource"
)

const (
	ImgNone resource.ImageID = iota
	ImgBackground
	ImgBtnIdle
	ImgBtnHover
	ImgBtnPressed
	ImgIconMenu
	ImgIconRestart
	ImgIconSubmit
	ImgFrame
	ImgInput
	ImgSliderTrack
	ImgSliderBtn
)

func RegisterImageResources(loader *resource.Loader) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImgBackground:  {Path: "background.png"},
		ImgBtnIdle:     {Path: "ui/btn-idle.png"},
		ImgBtnHover:    {Path: "ui/btn-hover.png"},
		ImgBtnPressed:  {Path: "ui/btn-pressed.png"},
		ImgIconMenu:    {Path: "ui/menu-icon.png"},
		ImgIconRestart: {Path: "ui/restart-icon.png"},
		ImgIconSubmit:  {Path: "ui/submit-icon.png"},
		ImgFrame:       {Path: "ui/frame.png"},
		ImgInput:       {Path: "ui/input.png"},
		ImgSliderTrack: {Path: "ui/slider-track.png"},
		ImgSliderBtn:   {Path: "ui/slider-btn.png"},
	}

	for id, res := range imageResources {
		loader.ImageRegistry.Set(id, res)
		loader.LoadImage(id)
	}
}

const (
	FontNone resource.FontID = iota
	FontDefault
)

func RegisterFontResources(loader *resource.Loader) {
	fontResources := map[resource.FontID]resource.FontInfo{
		FontDefault: {Path: "PrintChar21.ttf", Size: 6},
	}
	for id, res := range fontResources {
		loader.FontRegistry.Set(id, res)
		loader.LoadFont(id)
	}
}

func OpenAssetFunc(path string) io.ReadCloser {
	f, err := gameAssets.Open("resources/" + path)
	if err != nil {
		panic(err)
	}
	return f
}

//go:embed all:resources
var gameAssets embed.FS
