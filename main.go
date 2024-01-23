package main

import (
	"magic-gui/layout"
	"magic-gui/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	magicApp := app.New()
	magicTheme := theme.MyTheme{}
	magicApp.Settings().SetTheme(&magicTheme)
	magicWindow := magicApp.NewWindow("Magic Wechat")
	magicWindow.Resize(fyne.NewSize(800, 600))
	magicWindow.CenterOnScreen()

	layout.NewLayout(magicWindow)

	magicWindow.ShowAndRun()
}
