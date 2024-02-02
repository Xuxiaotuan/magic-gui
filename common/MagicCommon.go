package common

import "fyne.io/fyne/v2"

// MagicCommon defines the data structure for a MagicCommon
type MagicCommon struct {
	Title, Intro string
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

var (
	// MagicCommons defines the metadata for each MagicCommon
	MagicCommons = map[string]MagicCommon{
		"welcome":  {"首页", "", welcomeScreen, true},
		"download": {"下载", "", download, true},
	}

	// MagicCommonIndex  defines how our MagicCommons should be laid out in the index tree
	MagicCommonIndex = map[string][]string{
		"":            {"welcome", "download", "animations", "icons", "widgets", "collections", "containers", "dialogs", "windows", "binding", "advanced"},
		"collections": {"list", "table", "tree", "gridwrap"},
		"containers":  {"apptabs", "border", "box", "center", "doctabs", "grid", "scroll", "split"},
		"widgets":     {"accordion", "button", "card", "entry", "form", "input", "progress", "text", "toolbar"},
	}
)
