package view

import (
	"github.com/jmesyan/layerui"
	"github.com/andlabs/ui"
)

func box1()*layerui.LayContainer{
        box:=new(layerui.LayContainer)
	box.LcType=layerui.LcBox
	box.Stretchy=true
	box.Box=ui.NewVerticalBox()
	box.AddWidget("hbox",layerui.LayWidget{hbox1(),false})
	box.Compose()
	return box
}

func hbox1()*layerui.LayContainer{
	box:=new(layerui.LayContainer)
	box.LcType=layerui.LcBox
	box.Stretchy=false
	box.Box=ui.NewHorizontalBox()
	box.AddWidget("label",layerui.LayWidget{ui.NewLabel("counter"),true})
	box.AddWidget("counter",layerui.LayWidget{ui.NewEntry(),true})
	box.AddWidget("counter_plus",layerui.LayWidget{ui.NewButton("counter plus"),true})
	box.AddWidget("counter_minus",layerui.LayWidget{ui.NewButton("counter minus"),true})
	box.Compose()
	return box
}
