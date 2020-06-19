package main

import "fmt"

//Usb
type MUsb interface { //interface不能包含任何变量
	Start()
	Stop()
}

type MPhone struct {
	name string
}
func (p MPhone) Start() {
	fmt.Println("手机开始工作")
}
func (p MPhone) Stop() {
	fmt.Println("手机停止工作")
}
func (p MPhone) Call() {
	fmt.Println("手机在打电话")
}
//Phone
type MCamaera struct {
	name string
}
func (c MCamaera) Start() {
	fmt.Println("相机开始工作")
}
func (c MCamaera) Stop() {
	fmt.Println("相机停止工作")
}
//计算机
type MComputer struct {}
func (c MComputer) Working(usb MUsb){
	usb.Start()
	if mphone,ok := usb.(MPhone);ok {
		mphone.Call()
	}
	usb.Stop()
}

func main() {
	var usbArr [3]MUsb
	usbArr[0] = MPhone{"vivo"}
	usbArr[1] = MPhone{"小米"}
	usbArr[2] = MCamaera{"尼康"}
	
	var mcomp MComputer
	for _,v := range usbArr{
		mcomp.Working(v)
		fmt.Println()
	}
}
