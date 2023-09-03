package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"magic-gui/layout"
	"magic-gui/theme"
)

func main() {
	magicApp := app.New()
	magicTheme := theme.MyTheme{}
	magicApp.Settings().SetTheme(&magicTheme)
	magicWindow := magicApp.NewWindow("Magic Wechat")

	content := layout.NewLayout(magicWindow)
	magicWindow.SetContent(content)
	magicWindow.Resize(fyne.NewSize(800, 600))
	magicWindow.CenterOnScreen()

	magicWindow.ShowAndRun()
}
