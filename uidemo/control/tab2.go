package control

import (
	"github.com/jmesyan/layerui"
	"github.com/andlabs/ui"
	"github.com/jmesyan/layerui/uidemo/model"
)

func tab2Ctr(tab2 *layerui.LayContainer){
	box2:=tab2.Widget("box2").(*layerui.LayContainer)
	hbox:=box2.Widget("hbox").(*layerui.LayContainer)
	//get the model value
	msg:=model.GetMsgText()
	//button control
	buttonMsg:=hbox.Widget("msg_button").(*ui.Button)
	buttonMsg.OnClicked(func(b *ui.Button){
		w := ui.NewWindow("msg", 400, 300, false)
		ui.MsgBox(w, "msg tip", msg)
	})

}
