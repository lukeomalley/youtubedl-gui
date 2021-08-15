package main

import (
	_ "embed"
	"github.com/wailsapp/wails"
)

//go:embed frontend/public/build/bundle.js
var js string

//go:embed frontend/public/build/bundle.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "ytdl-frontend",
		JS:        js,
		CSS:       css,
		Colour:    "#16191e",
		Resizable: true,
	})

	app.Bind(NewDownloader())
	app.Run()
}
