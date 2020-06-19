package main

import (
	"fmt"
	"os"
)

//获取错误的类型
func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		//错误类型表示
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("1、Op：", ins.Op)
			fmt.Println("2、Path：", ins.Path)
			fmt.Println("3、Err：", ins.Err)
		}
		return
	}
	fmt.Println("打开文件成功", file)
}
