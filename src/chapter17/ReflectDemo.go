package main

import (
	"Go-Player/src/chapter17/models"
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	//通过反射获取到传入变量的type，kind，值
	
	//1、先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType:=", rType) //rType:= int
	
	//2、获取到reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("值=%d，类型=%T\n", rValue, rValue) //值=100，类型=reflect.Value
	
	//注意：它的rType虽然输出int，但是并不是实际意义上的int
	//需要写成如下形式
	num2 := 10 + rValue.Int() //num2= 110
	fmt.Println("num2=", num2)

	//3、将rValue转化成interface{}
	iV := rValue.Interface()
	
	//4、将interface{}通过断言转化成需要的类型
	num3 := iV.(int)
	fmt.Printf("值=%d，类型=%T\n", num3, num3) //值=100，类型=int
}

func reflectTest02(b interface{}) {
	//通过反射获取传入的变量的type，kind值
	
	//1、先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType:=", rType) //rType:= models.Student
	
	//2、获取到reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("值=%v，类型=%T\n", rValue, rValue) //值={张三 19}，类型=reflect.Value
	
	//3、将rVal转化成interface{}
	iV := rValue.Interface()
	
	//4、将interface{}通过断言转化成相应的形式
	//使用待检测的类型断言
	//可以使用switch断言，使内容更加灵活
	/*switch iV.(type) {
	case models.Student:
		fmt.Println("Student")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("判断不出来类型")
	}*/
	stu, ok := iV.(models.Student)
	if ok {
		fmt.Printf("stu.Name = %v\n", stu.Name) //stu.Name = 张三
	}
}

func main() {
	var num int = 100
	reflectTest01(num)
	
	stu := models.Student{
		Name: "张三",
		Age:  19,
	}
	reflectTest02(stu)
}
