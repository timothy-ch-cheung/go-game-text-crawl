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
	ImgFrame
)

func RegisterImageResources(loader *resource.Loader) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImgBackground:  {Path: "background.png"},
		ImgBtnIdle:     {Path: "ui/btn-idle.png"},
		ImgBtnHover:    {Path: "ui/btn-hover.png"},
		ImgBtnPressed:  {Path: "ui/btn-pressed.png"},
		ImgIconMenu:    {Path: "ui/menu-icon.png"},
		ImgIconRestart: {Path: "ui/restart-icon.png"},
		ImgFrame:       {Path: "ui/frame.png"},
	}

	for id, res := range imageResources {
		loader.ImageRegistry.Set(id, res)
		loader.LoadImage(id)
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
