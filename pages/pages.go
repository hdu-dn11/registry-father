package pages

import (
	"github.com/rivo/tview"
	"registry-father/pages/asManage"
)

func Load(app *tview.Application) tview.Primitive {
	return tview.NewPages().
		AddPage("AsManage", asManage.Load(app), true, true)
}
