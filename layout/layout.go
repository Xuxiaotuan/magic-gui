package layout

import (
	"magic-gui/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	content = container.NewStack()
)

func NewLayout(window fyne.Window) {
	sourceInput := widget.NewMultiLineEntry()
	filePathInput := widget.NewMultiLineEntry()
	keyOutput := widget.NewEntry()
	filePathOutput := widget.NewEntry()
	generateButton := widgets.GenerateButton(sourceInput, keyOutput)
	progress := widget.NewProgressBar()
	runButton := widgets.RunButton(progress, sourceInput, filePathInput)

	navgater := container.NewVBox(
		widget.NewLabel("请输入待解析的数据,举例 \n"+
			"0x600007936780: 0xdc 0x8a 0xb2 0x78 0x9a 0x66 0x41 0x79   \n"+
			"0x600007936788: 0xad 0x34 0xd1 0xb2 0x3b 0x82 0x73 0x89   \n"+
			"0x600007936790: 0xc6 0x31 0x8a 0xbe 0x2d 0x62 0x4f 0x05   \n"+
			"0x600007936798: 0xa7 0xac 0xb7 0xf7 0xa9 0x17 0x36 0xe0"),
		sourceInput,
		generateButton,
		widget.NewLabel("生成的微信key:"),
		keyOutput,
		widget.NewLabel("请输入电脑中的路径:"),
		filePathInput,
		widget.NewLabel("请输入解密文件后存放的路径:"),
		filePathOutput,
		runButton,
		progress,
	)
	page := container.NewHSplit(navgater, content)
	page.SetOffset(0.3)

	window.SetContent(page)
}
