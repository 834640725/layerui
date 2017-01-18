package view

import (
	"github.com/jmesyan/layerui"
	"github.com/andlabs/ui"
)

func box2()*layerui.LayContainer{
	box:=new(layerui.LayContainer)
	box.LcType=layerui.LcBox
	box.Stretchy=true
	box.Box=ui.NewVerticalBox()
	box.AddWidget("hbox",layerui.LayWidget{hbox2(),false})
	box.Compose()
	return box
}

func hbox2()*layerui.LayContainer{
	box:=new(layerui.LayContainer)
	box.LcType=layerui.LcBox
	box.Stretchy=false
	box.Box=ui.NewHorizontalBox()
	box.AddWidget("label",layerui.LayWidget{ui.NewLabel("msg button"),true})
	box.AddWidget("msg_button",layerui.LayWidget{ui.NewButton("click me"),true})
	box.Compose()
	return box
}
