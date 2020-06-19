package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var b bool
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type is : %T, b = %v\n", b, b) //b type is : bool, b = true
	
	str = "123"
	i, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("i type is : %T, i= %v\n",i, i) //i type is : int64, i= 123
	
	str = "123.456"
	f, _ := strconv.ParseFloat(str, 64)
	fmt.Printf("f type is : %T, f= %v\n",f, f) //f type is : float64, f= 123.456
	
	str = "hello"
	i2, _ := strconv.ParseInt(str, 10, 64)//这列str不能转化成数字，转化失败的话会变成0值或者默认值
	fmt.Printf("i2 type is : %T, i2= %v\n",i2, i2) //i2 type is : int64, i2= 0
	
}
