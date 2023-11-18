package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"changeme/pkg/api"
	etcd2 "changeme/pkg/etcd"
	"changeme/pkg/global"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	if err := global.Init(); err != nil {
		panic(err)
	}

	etcdApi := api.NewEtcdApi()
	etcdClient := etcd2.MustGetEtcdClient()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "hci-etcd",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		// 指定向前端暴露哪些结构体方法
		Bind: []interface{}{
			app,
			etcdApi,
			etcdClient,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
