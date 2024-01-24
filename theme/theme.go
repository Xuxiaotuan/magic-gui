package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"

	"image/color"
)

type XxtTheme struct{}

var _ fyne.Theme = (*XxtTheme)(nil)

// return bundled font resource
// resourceFontTtf 即是 bundle.go 文件中 var 的变量名
func (m XxtTheme) Font(s fyne.TextStyle) fyne.Resource {
	return resourceFontTtf
}
func (*XxtTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*XxtTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*XxtTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
