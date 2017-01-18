package control

import (
	"github.com/jmesyan/layerui"
	"github.com/andlabs/ui"
	"github.com/jmesyan/layerui/uidemo/model"
	"strconv"
)

func tab1Ctr(tab1 *layerui.LayContainer){
	box1:=tab1.Widget("box1").(*layerui.LayContainer)
	hbox:=box1.Widget("hbox").(*layerui.LayContainer)
	//init the counter
	counter:=hbox.Widget("counter").(*ui.Entry)
	counter.SetReadOnly(true)
	//get the model value
	count:=model.GetDefaultCount()
	counter.SetText(count)
	//button control
	buttonPlus:=hbox.Widget("counter_plus").(*ui.Button)
	buttonPlus.OnClicked(func(b *ui.Button){
		textNum,_:=strconv.Atoi(counter.Text())
		num:=textNum+1
		counter.SetText(strconv.Itoa(num))
	})

	buttonMinus:=hbox.Widget("counter_minus").(*ui.Button)
	buttonMinus.OnClicked(func(b *ui.Button){
		textNum,_:=strconv.Atoi(counter.Text())
		num:=textNum-1
		counter.SetText(strconv.Itoa(num))
	})

}
