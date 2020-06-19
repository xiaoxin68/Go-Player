package main

import "fmt"

func main() {
	var sliceOfMap []map[string]string
	sliceOfMap = make([]map[string]string, 2) //sliceOfMap是切片，需要make一下才能使用
	if sliceOfMap[0] == nil {
		sliceOfMap[0] = make(map[string]string, 2) //map也需要make一下才能使用
		sliceOfMap[0]["name"] = "牛魔王"
		sliceOfMap[0]["address"] = "火焰山"
	}
	if sliceOfMap[1] == nil {
		sliceOfMap[1] = map[string]string{
			"name":    "孙悟空",
			"teacher": "唐僧",
			"address": "花果山",
		}
	}
	
	fmt.Println(sliceOfMap)
	
	//由于slice可以动态扩展
	newMap := make(map[string]string, 2)
	newMap["name"] = "观音菩萨"
	newMap["age"] = "一千岁"
	
	//追加newMap到sliceOfMap
	sliceOfMap = append(sliceOfMap, newMap)
	fmt.Println(sliceOfMap)
	
}
