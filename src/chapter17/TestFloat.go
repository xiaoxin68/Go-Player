package main

import (
	"fmt"
	"reflect"
)

func testFloat(b interface{})  {
	//获取变量的val
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)
	
	//获取变量的Type
	rType := reflect.TypeOf(b)
	fmt.Println(rType)
	
	//获取变量的Kind
	kind1 := rVal.Kind()
	kind2 := rType.Kind()
	fmt.Println(kind1," ",kind2)
	
	//将val转化成interface{}
	num := rVal.Interface()
	fmt.Println(num)
	
	//将interface{}转化成float
	res := num.(float64)
	fmt.Println(res)
}
func main() {
	var num float64 = 12.34
	testFloat(num)
	fmt.Printf("num=%v\n",num)
}
