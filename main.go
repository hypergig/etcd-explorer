package main

import (
	"embed"
	"github.com/hypergig/etcd-explorer/internal/listwatcher"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	//app := NewApp()
	service := listwatcher.New()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "etcd-explorer",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        service.SetCtx,
		Bind: []interface{}{
			service,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
