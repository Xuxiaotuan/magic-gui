package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// MagicButton 继承widget.Button用于个性化定制
type MagicButton struct {
	widget.Button
}

// Tapped 重写Tapped方法, 捕获点击事件的同时,获取button的属性信息并执行一系列操作
func (b *MagicButton) Tapped(*fyne.PointEvent) {

	if b.Disabled() {
		return
	}
	b.Text = "正在运行中"

	if b.OnTapped != nil {
		b.OnTapped()
		b.Disable()
	}
	b.Refresh()

}

// NewButton MagicButton的工厂方法
func XxtButton(label string, tapped func()) *MagicButton {
	button := &MagicButton{}
	button.Text = label
	button.OnTapped = tapped
	button.Alignment = 0

	button.ExtendBaseWidget(button)
	return button
}
