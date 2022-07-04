package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"notepage.de/desktop/data"
)

func main() {
	a := app.New()
	win := a.NewWindow("Hello World")
	fontColor := color.White
	themeType := theme.LightTheme()

	// Setting theme from default config
	if data.THEME == "light" {
		themeType = theme.LightTheme()
		fontColor = color.Black
	} else if data.THEME == "dark" {
		themeType = theme.DarkTheme()
		fontColor = color.White
	}

	a.Settings().SetTheme(themeType)

	// Set on Screen values
	title := canvas.NewText(data.TITLE, fontColor)
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	win.Resize(fyne.NewSize(float32(data.WIDTH), float32(data.HEIGHT)))

	hBox := container.New(layout.NewHBoxLayout())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox)

	win.SetContent(vBox)
	win.ShowAndRun()
}
