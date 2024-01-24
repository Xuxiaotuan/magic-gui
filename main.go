package main

import (
	"magic-gui/layout"
	"magic-gui/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	magicApp := app.New()
	magicTheme := theme.XxtTheme{}
	magicApp.Settings().SetTheme(&magicTheme)
	magicWindow := magicApp.NewWindow("Magic Downloads")
	magicWindow.Resize(fyne.NewSize(800, 600))
	magicWindow.CenterOnScreen()

	// 设置布局
	layout.NewLayout(magicWindow)

	// 常驻服务
	magicWindow.ShowAndRun()
}
