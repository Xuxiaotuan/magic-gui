package widgets

import (
	"fyne.io/fyne/v2/widget"
	"strings"
	"time"
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

func RunButton(progress *widget.ProgressBar) *MyButton {
	button := NewButton("点我开始进行解析同步", func() {
		go func() {
			progress.SetValue(0.0)
			for i := 0.0; i <= 1.0; i += 0.1 {
				time.Sleep(time.Millisecond * 250)
				progress.SetValue(i)
			}
		}()
	})
	return button
}
