package widgets

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2/widget"
)

func GenerateButton(sourceInput *widget.Entry, keyOutput *widget.Entry) *widget.Button {
	return widget.NewButton("点我生成密钥", func() {
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
}

func RunButton(progress *widget.ProgressBar, input *widget.Entry, pathInput *widget.Entry) *MyButton {
	return NewButton("点我开始进行解析同步", func() {
		go func() {
			if len(input.Text) == 0 {
				fmt.Println("电脑中的文件路径为空")
			}
			if len(pathInput.Text) == 0 {
				fmt.Println("解密后文件存放的路径为空")
			}
			fmt.Printf("%s", input.Text)
			fmt.Printf("%s", pathInput.Text)
			progress.SetValue(0.0)
			for i := 0.0; i <= 1.0; i += 0.1 {
				time.Sleep(time.Millisecond * 250)
				progress.SetValue(i)
			}
		}()
	})
}
