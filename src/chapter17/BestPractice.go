package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

//返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//接收四个值，给s赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

//显示s的值
func (s Monster) Print() {
	fmt.Println("----start----")
	fmt.Println(s)
	fmt.Println("----end----")
}

//使用反射获取结构体内容
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	
	//变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Filed %d：值为=%v\n", i, val.Field(i))
		//获取到struct标签：注意需要通过reflect.Type来获取tag
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段属于tag标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Filed %d:tag 为=%v\n", i, tagVal)
		}
	}
	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	
	//var  params []reflect.Value
	//方法的默认顺序是按照函数名的排序
	val.Method(1).Call(nil) //获取到第二个方法，调用它
	
	//调用结构体的第一个方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //Call方法的参数是 []reflect.Value，返回 []reflect.Value
	fmt.Println("res=", res[0].Int()) //返回的结果是[]reflect.Value
}

//修改值
func UpdateStruct(a interface{}) {
	val := reflect.ValueOf(a)
	val.Elem().Field(0).SetString("白象精")
}
func main() {
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	TestStruct(a)
	
	UpdateStruct(&a)
	fmt.Println(a)
}
