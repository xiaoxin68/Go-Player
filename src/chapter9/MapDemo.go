package main

import "fmt"

//map中嵌套map
func main() {
	//key:string
	//val:map[string]string
	students := make(map[string]map[string]string, 2)
	
	students["stu01"] = make(map[string]string, 3)
	students["stu01"]["name"] = "张三"
	students["stu01"]["age"] = "20岁"
	students["stu01"]["address"] = "湖北武汉"
	
	students["stu02"] = make(map[string]string, 3)
	students["stu02"]["name"] = "李四"
	students["stu02"]["age"] = "21岁"
	students["stu02"]["address"] = "湖北随州"
	
	fmt.Println(students)
	fmt.Println(students["stu01"])
	fmt.Println(students["stu02"]["address"])
}
