package main

import "fmt"

func main() {
	test1()
	fmt.Println("main end")
}

func test1()  {
	//使用defer和recover来捕获和处理异常
	defer func() {
		err := recover()  //捕获异常
		if err != nil{  //说明捕获到错误
			fmt.Println("err=",err)
		}
	}()
	
	fmt.Println("test1 enter")
	var num1,num2  int = 10,0
	fmt.Println(num1/num2)
	fmt.Println("test1 end")
}

//运行结果：
//test1 enter
//err= runtime error: integer divide by zero
//main end
