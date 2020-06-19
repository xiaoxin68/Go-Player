package main

import "fmt"

func main()  {
	fmt.Println("张三\t李四")
	fmt.Println("张三\n李四")
	fmt.Println("张三\\李四")
	fmt.Println("张三\"李四")
	fmt.Println("张三\r李四")  //表示从当行的最前面输出，替换掉前面的内容
}
