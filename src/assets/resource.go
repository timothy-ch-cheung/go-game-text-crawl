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
	ImgMenuBtnIdle
	ImgMenuBtnHover
	ImgMenuBtnPressed
	ImgRestartBtnIdle
	ImgRestartBtnHover
	ImgRestartBtnPressed
	ImgFrame
)

func RegisterImageResources(loader *resource.Loader) {
	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImgBackground:        {Path: "background.png"},
		ImgMenuBtnIdle:       {Path: "ui/menu-btn-idle.png"},
		ImgMenuBtnHover:      {Path: "ui/menu-btn-hover.png"},
		ImgMenuBtnPressed:    {Path: "ui/menu-btn-pressed.png"},
		ImgRestartBtnIdle:    {Path: "ui/restart-btn-idle.png"},
		ImgRestartBtnHover:   {Path: "ui/restart-btn-hover.png"},
		ImgRestartBtnPressed: {Path: "ui/restart-btn-pressed.png"},
		ImgFrame:             {Path: "ui/frame.png"},
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
