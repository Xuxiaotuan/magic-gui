package widgets

import (
	"fyne.io/fyne/v2/widget"
	"strings"
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

func RunButton(sourceInput *widget.Entry, keyOutput *widget.Entry) *widget.Button {
	return widget.NewButton("点我开始进行解析同步", func() {
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
