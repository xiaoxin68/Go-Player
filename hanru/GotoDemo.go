package main

import "fmt"

//go语言return语句后面居然还可以有内容
//goto语句可以方便进行集中错误处理
func main() {
	for i:=0;i<10;i++ {
		if i==5 {
			 goto breakFlag
		}
	}
	//注意这里要加一个return语句，避免程序执行了标签后面的内容
	return

	fmt.Println("hhahahah")

	breakFlag:
		fmt.Println("这里")
}
