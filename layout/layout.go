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
	// 下载视频
	sourceInput := widget.NewMultiLineEntry()
	ytlpPathInput := widget.NewMultiLineEntry()
	outVideoPathInput := widget.NewMultiLineEntry()
	generateButton := widgets.GenerateButton(sourceInput, ytlpPathInput, outVideoPathInput)

	// 转换视频
	ffmpegResultPathInput := widget.NewMultiLineEntry()
	filePathOutput := widget.NewEntry()
	runButton := widgets.RunButton(window, outVideoPathInput, ffmpegResultPathInput)

	navgater := container.NewVBox(
		widget.NewLabel("需要输入要下载的视频链接，比如 \n"+
			"https://www.youtube.com/watch?v=jgVhBThJdXc"),
		sourceInput,
		widget.NewLabel("需要手动下载yt-dlp的文件，并给出文件路径:"),
		ytlpPathInput,
		widget.NewLabel("转换后视频想变成什么样呢： 比如： /Users/xjw/software/可惜我是水瓶座伴奏【杨千嬅】 [BV1Fv4y1c7in_p1].mp4"),
		outVideoPathInput,
		generateButton,
		widget.NewLabel("请输入转换后文件存放的路径:"),
		filePathOutput,
		runButton,
	)
	page := container.NewHSplit(navgater, content)
	page.SetOffset(0.3)

	window.SetContent(page)
}
