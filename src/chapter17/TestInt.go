package main

import (
	"fmt"
	"reflect"
)

//通过反射修改numInt的值
func testInt(b interface{}) {
	val := reflect.ValueOf(b)
	fmt.Printf("val kind=%v\n", val.Kind()) //val kind=ptr
	val.Elem().SetInt(110)
	fmt.Printf("val type=%T\n", val) //val type=reflect.Value
}

func main() {
	var num int = 20
	testInt(&num)            //注意这里传的是地址
	fmt.Println("num=", num) //num= 110
}

//val.Elem()用于获取指针指向变量，类似于
/*var num = 10
var ptr *int = &num
num2 := *ptr  //===类似val.Elem()*/