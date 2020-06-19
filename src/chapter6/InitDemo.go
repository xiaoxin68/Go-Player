package main

import "fmt"

//最先完成全局变量初始化
//其次完成init
//最后完成main
/*执行结果：
test...
Init....
Main...*/

var a = test2()

func test2()int  {
	fmt.Println("test...")
	return 10
}

//init函数会在main函数之前执行 ：主要用于完成初始化工作
func init() {
	fmt.Println("Init....")
}

func main() {
	fmt.Println("Main...")
}
