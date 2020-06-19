package main

//json序列化
//func Marshal(v interface{}) ([]byte, error):Marshal函数返回v的json编码。

import (
	"encoding/json"
	"fmt"
	"./models"
)

//对结构体进行序列化
func testDtruct() {
	//定义初始化结构体
	monster := models.Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "1998-10-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}
	
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
}

//对map进行序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	//使用map需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"
	//map序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("a map 序列化后=%v\n", string(data))
}

//切片序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	//使用map前，先make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = "北京"
	slice = append(slice, m1)
	
	var m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 20
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)
	
	var m3 = make(map[string]interface{})
	m3["monster"] = models.Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "1998-10-11",
		Sal:      8000.0,
		Skill:    "牛魔拳",
	}
	slice = append(slice, m3)
	
	//序列化切片
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("slice 序列化后=%v\n", string(data))
}

//对基本数据类型序列化：对基本数据类型进行序列化意义不大
func testFloat64()  {
	var num1 float64 = 2345.67
	//对num1序列化
	data,err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("slice 序列化后=%v\n", string(data))
}
func main() {
	testDtruct()
	testMap()
	testSlice()
	testFloat64()
}
