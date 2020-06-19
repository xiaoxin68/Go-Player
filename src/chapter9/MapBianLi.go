package main

import "fmt"

func main() {
	var a = make(map[string]string, 10)
	a["idx1"] = "hello"
	a["idx2"] = "world"
	a["idx3"] = "zxx"
	
	//简单map遍历
	for key, val := range a {
		fmt.Printf("key=%s,val=%s\n", key, val)
	}
	
	var stuMap = make(map[string]map[string]string)
	stuMap["stu01"] = make(map[string]string, 3)
	stuMap["stu01"]["name"] = "张三"
	stuMap["stu01"]["age"] = "20岁"
	stuMap["stu01"]["address"] = "湖北武汉"
	
	stuMap["stu02"] = map[string]string{
		"name":    "李四",
		"age":     "21岁",
		"address": "湖北随州",
	}
	
	//复杂map的遍历
	for key,val := range stuMap{
		fmt.Printf("outerKey=%s\n", key)
		for key2,val2 := range val {
			fmt.Printf("\tinnerKey=%s,val=%s\n",key2,val2)
		}
	}
}
