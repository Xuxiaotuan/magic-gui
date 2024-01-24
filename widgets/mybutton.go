package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// MyButton 继承widget.Button用于个性化定制
type MyButton struct {
	widget.Button
}

// Tapped 重写Tapped方法, 捕获点击事件的同时,获取button的属性信息并执行一系列操作
func (b *MyButton) Tapped(*fyne.PointEvent) {

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

// NewButton MyButton的工厂方法
func NewButton(label string, tapped func()) *MyButton {
	button := &MyButton{}
	button.Text = label
	button.OnTapped = tapped
	button.Alignment = 1

	button.ExtendBaseWidget(button)
	return button
}
