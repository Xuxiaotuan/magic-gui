package widgets

import (
	"fmt"
	"magic-gui/util"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func GenerateButton(sourceInput *widget.Entry, ytlpPathInput *widget.Entry, resultPathInput *widget.Entry) *widget.Button {
	return widget.NewButton("点我下载视频～", func() {
		util.DownloadViedo(ytlpPathInput.Text, sourceInput.Text, resultPathInput.Text)
	})
}

func confirmCallback(response bool) {
	fmt.Println("Responded with", response)
}

func RunButton(window fyne.Window, input *widget.Entry, pathInput *widget.Entry) *MyButton {
	return NewButton("点我开始进行转换", func() {
		go func() {
			if len(input.Text) == 0 || len(pathInput.Text) == 0 {
				cnf := dialog.NewConfirm("Confirmation", "电脑中的视频 / 转换后文件存放的路径 为空", confirmCallback, window)
				cnf.SetDismissText("返回")
				cnf.SetConfirmText("确认")
				cnf.Show()
			} else {
				util.CovertMp4ToMp3(input.Text, pathInput.Text)
			}
		}()
	})
}
