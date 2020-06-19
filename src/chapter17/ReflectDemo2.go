package main

import (
	"fmt"
	"reflect"
)

//两种方式获取变量对应的Kind
func getReflectKind(b interface{}) {
	rType := reflect.TypeOf(b)
	
	rVal := reflect.ValueOf(b)
	
	fmt.Println(rType.Kind()) //int
	
	fmt.Println(rVal.Kind()) //int
}
func main() {
	num := 10
	getReflectKind(num)
}
