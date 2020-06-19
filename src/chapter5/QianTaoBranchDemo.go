package main

import "fmt"

func main() {
	var seconds float32
	fmt.Println("请输入时间：")
	fmt.Scanln(&seconds)
	
	if seconds <= 8 {
		var gender string
		fmt.Println("请输入性别：")
		fmt.Scanf("%s", &gender)
		if gender == "男" {
			fmt.Println("进入男子组")
		} else {
			fmt.Println("进入女子组")
		}
	} else {
		fmt.Println("out")
	}
}
