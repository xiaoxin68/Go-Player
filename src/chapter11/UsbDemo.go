package main

import "fmt"

//Usb
type Usb interface { //interface不能包含任何变量
	Start()
	Stop()
}
//Phone：一个自定义类型可以实现多个接口
//接口不需要显式实现，只要一个变量，含有接口类型中的所有方法，那么这个变量就实现这个接口
type Phone struct {}
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}
//Phone
type Camaera struct {}
func (c Camaera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camaera) Stop() {
	fmt.Println("相机停止工作")
}
//计算机
type Computer struct {}
func (c Computer) Working(usb Usb){
	usb.Start()
	usb.Stop()
}

//只要是自定义类型就可以实现接口，不仅仅是结构体类型
type MyInteger int
func (i MyInteger)Start()  {
	fmt.Println("自定义【非结构体类型】start")
}
func (i MyInteger)Stop()  {
	fmt.Println("自定义【非结构体类型】stop")
}
func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camaera{}
	
	computer.Working(phone) //注意：接口是引用类型，接收的是地址
	computer.Working(camera)
	
	fmt.Println()
	
	//接口不能创建实例变量,但是可以指向一个实现了接口的自定义类型
	var usb Usb = Phone{}
	usb.Start()
	var cam Camaera
	usb1 := cam
	usb1.Start()
	
	fmt.Println()
	
	//只要是自定义类型就可以实现接口，不仅仅是结构体类型
	var i  MyInteger = 10
	usd2 := i
	computer.Working(i)
	fmt.Println(usd2)
	
	//多态
	var usbs [3]Usb
	usbs[0] = Phone{}
	usbs[1] = Camaera{}
	usbs[2] = Camaera{}
	for idx,val := range usbs {
		fmt.Printf("usb[%d]的类型为:%T\n",idx+1,val)
	}
}
