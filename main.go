package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"magic-gui/theme"
	"strings"
)

func main() {
	myApp := app.New()
	myTheme := theme.MyTheme{}
	myTheme.SetFonts("./data/font.ttf", "")
	myApp.Settings().SetTheme(&myTheme)
	myWindow := myApp.NewWindow("Magic Wechat")

	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(800, 500))

	sourceInput := widget.NewMultiLineEntry()
	filePathInput := widget.NewMultiLineEntry()
	keyOutput := widget.NewEntry()
	filePathOutput := widget.NewEntry()

	generateButton := widget.NewButton("Generate Key", func() {
		source := sourceInput.Text
		var data []string
		if source != "" {
			for _, line := range strings.Split(source, "\n") {
				fields := strings.Fields(line)
				for _, field := range fields[1:] {
					data = append(data, strings.ReplaceAll(field, "0x", ""))
				}
			}
		}
		key := "0x" + strings.Join(data, "")
		keyOutput.SetText(key)
	})

	content := container.NewVBox(
		widget.NewLabel("请输入待解析的数据"),
		sourceInput,
		generateButton,
		widget.NewLabel("生成的微信key:"),
		keyOutput,
		widget.NewLabel("请输入电脑中的路径:"),
		filePathInput,
		widget.NewLabel("请输入解密后的路径:"),
		filePathOutput,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
