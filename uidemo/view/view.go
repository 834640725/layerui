package view

import (
	"github.com/jmesyan/layerui"
	"github.com/andlabs/ui"
)

func tab1()*layerui.LayContainer{
	box:=new(layerui.LayContainer)
	box.LcType=layerui.LcBox
	box.Title="box1_title"
	box.Stretchy=true
	box.Box=ui.NewVerticalBox()
	box.AddWidget("box1",layerui.LayWidget{box1(),true})
	box.Compose()
	return box
}


func tab2()*layerui.LayContainer{
	box:=new(layerui.LayContainer)
	box.LcType=layerui.LcBox
	box.Title="box2_title"
	box.Stretchy=true
	box.Box=ui.NewVerticalBox()
	box.AddWidget("box2",layerui.LayWidget{box2(),true})
	box.Compose()
	return box
}

func MainTab()*layerui.LayContainer{
	tab:=new(layerui.LayContainer)
	tab.LcType=layerui.LcTab
	tab.Tab=ui.NewTab()
	tab.AddWidget("tab1",layerui.LayWidget{tab1(),true})
	tab.AddWidget("tab2",layerui.LayWidget{tab2(),true})
	tab.Compose()
	return tab
}