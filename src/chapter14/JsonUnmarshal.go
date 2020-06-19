package main

//json反序列化
//func Unmarshal(data []byte, v interface{}) error ：Unmarshal函数解析json编码的数据并将结果存入v指向的值。

import (
	"./models"
	"encoding/json"
	"fmt"
)

//反将json字符串反序列化成结构体
func unmarshalStruct() {
	//注意：如果json字符串是通过程序获取到的，就不需要再对双引号进行转义处理
	str := "{\"monster_name\":\"牛魔王\",\"monster_age\":500,\"Birthday\":\"1998-10-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
	var monster models.Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("反序列化后monster=%v monster.name=%v\n", monster, monster.Name)
}

//将json字符串反序列化成map
func unmarshalMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("反序列化后a map=%v \n", a)
}

//将json字符串反序列化成切片
func unmarshalSlice() {
	str := "[" +
		"{\"address\":\"北京\",\"age\":7,\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\",\"夏威夷\"],\"age\":20,\"name\":\"tom\"}," +
		"{\"monster\":{\"monster_name\":\"牛魔王\",\"monster_age\":50,\"Birthday\":\"1998-10-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}}" +
		"]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str),&slice)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("反序列化后slice=%v \n", slice)
}

//将json字符串反序列化成基本类型
func unmarshalFloat64()  {
	str := "2345.67"
	var num float64
	err := json.Unmarshal([]byte(str), &num)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("slice 序列化后=%v\n", num)
}
func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
	unmarshalFloat64()
}
