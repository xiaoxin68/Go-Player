package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
}
//给monster绑定方法Store，可以将一个Monster变量序列化后保存到文件中
func (this *Monster)Store()bool  {
	//先序列化
	data, err := json.Marshal(this)
	if err!=nil {
		fmt.Println("marshal err=",err)
		return false
	}
	
	//保存到文件
	filePath := "C:/Users/zxx/Desktop/a.txt"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err!=nil {
		fmt.Println("write file err=",err)
		return false
	}
	return true
}

//给monster绑定方法ReStore，可以将一个序列化的Monster从文件中读取，并反序列化成MMonster对象
func (this *Monster)ReStore()bool  {
	//读取文件
	filePath := "C:/Users/zxx/Desktop/a.txt"
	data, err := ioutil.ReadFile(filePath)
	if err!=nil {
		fmt.Println("read file err=",err)
		return false
	}
	
	//将读取到的文件进行反序列化
	err = json.Unmarshal(data, this)
	if err!=nil {
		fmt.Println("Unmarshal err=",err)
		return false
	}
	return true
}
func main() {
	
}
