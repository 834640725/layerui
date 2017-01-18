package control

import (
	"github.com/andlabs/ui"
	"github.com/jmesyan/layerui/uidemo/view"
	"github.com/jmesyan/layerui"
)

func UiCtr(){
	err := ui.Main(func() {
		//organize the ui
		MainTab :=view.MainTab()
		//control the ui
		//tab1
		tab1:=MainTab.Widget("tab1").(*layerui.LayContainer)
		tab1Ctr(tab1)
		//tab2
		tab2:=MainTab.Widget("tab2").(*layerui.LayContainer)
		tab2Ctr(tab2)
		//finally run the main window
		var Window = ui.NewWindow("window_title", 1000, 800, false)
		Window.SetChild(MainTab.Tab)
		Window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		Window.Show()
	})

	if err != nil {
		panic(err)
	}
}
