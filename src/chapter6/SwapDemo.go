package main

import "fmt"

func main()  {
	var num1,num2 int = 10,20
	swap(&num1,&num2)
	fmt.Printf("num1=%d,num2=%d",num1,num2)
}
func swap(num1 *int,num2 *int)  {
	temp := *num1
	*num1 = *num2
	*num2 = temp
}

//不能这么定义，go不支持函数的重载
//func swap()  {
//	fmt.Printf("kalla")
//}