package main

import "fmt"

var Age int = 10 //ok
/*错误: Name := "zhang"
相当于var Name string
Name = "zhang"
而赋值语句不能卸载函数体外！！！！
*/
//Name := "zhang"
var Name string = "zhang"

func main()  {
	fmt.Println(Age)
	fmt.Println(Name)
}
