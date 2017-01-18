package layerui

import "github.com/andlabs/ui"

const (
	LcTab=iota
	LcBox
)

type LayWidget struct {
	Widget   interface{}
	Stretchy bool
}

type LayContainer struct {
	LcType int
	Box   *ui.Box
	Tab   *ui.Tab
	Title    string
	Stretchy bool
	LywSort    []string
	LyWidgets  map[string]LayWidget
}

func (la *LayContainer)AddWidget(name string,widget LayWidget){
	if(la.LyWidgets==nil){
		la.LyWidgets=make(map[string]LayWidget)
	}
	la.LywSort=append(la.LywSort,name)
	la.LyWidgets[name]=widget
}

func (la *LayContainer)Widget(name string) interface{} {
	w := la.LyWidgets[name]
	return w.Widget
}

func (la *LayContainer)Compose() {

	if la.LcType==LcTab{
		tab:=la.Tab
		for _,k:=range la.LywSort{
			w:=la.LyWidgets[k]
			switch wig := w.Widget.(type){
			case *ui.Label:
				tab.Append(wig.Text(),wig)
			case *ui.Button:
				tab.Append(wig.Text(),wig)
			case *LayContainer:
				if wig.LcType==LcTab{
					tab.Append(wig.Title,wig.Tab)
				}else{
					tab.Append(wig.Title,wig.Box)
				}

			}
		}
	}else{
		box:=la.Box
		for _,k:=range la.LywSort{
			w:=la.LyWidgets[k]
			switch wig := w.Widget.(type){
			case *ui.Label:
				box.Append(wig, w.Stretchy)
			case *ui.Button:
				box.Append(wig, w.Stretchy)
			case *ui.Combobox:
				box.Append(wig,w.Stretchy)
			case *ui.Entry:
				box.Append(wig,w.Stretchy)
			case *ui.DateTimePicker:
				box.Append(wig,w.Stretchy)
			case *ui.Checkbox:
				box.Append(wig,w.Stretchy)
			case *LayContainer:
				if wig.LcType==LcTab{
					box.Append(wig.Tab,wig.Stretchy)
				}else{
					box.Append(wig.Box,wig.Stretchy)
				}
			}
		}

	}
}
