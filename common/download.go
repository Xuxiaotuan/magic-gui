package common

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"magic-gui/widgets"
)

func download(window fyne.Window) fyne.CanvasObject {
	// 下载视频
	sourceInput := widget.NewMultiLineEntry()
	ytlpPathInput := widget.NewEntry()
	outVideoPathInput := widget.NewMultiLineEntry()
	generateButton := widgets.GenerateButton(sourceInput, ytlpPathInput, outVideoPathInput)

	// 转换视频
	ffmpegResultPathInput := widget.NewMultiLineEntry()
	runButton := widgets.RunButton(window, outVideoPathInput, ffmpegResultPathInput)

	return container.NewVBox(
		widget.NewLabel("需要输入要下载的视频链接，比如: \n"+
			"https://www.youtube.com/watch?v=jgVhBThJdXc"),
		sourceInput,
		widget.NewLabel("需要手动下载yt-dlp的文件，并给出文件路径:"),
		ytlpPathInput,
		widget.NewLabel("转换后视频想变成什么样呢： 比如: \n"+
			"/Users/xjw/software/111.mp4"),
		outVideoPathInput,
		generateButton,
		widget.NewLabel("需要本地有ffmpeg  \n"+
			"请输入转换后文件存放的路径: 比如: \n"+
			"/Users/xjw/software/111.mp3"),
		ffmpegResultPathInput,
		runButton,
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	)
}
