package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	//file, err := os.Open("a_copy.txt") //open a_copy.txt: The system cannot find the file specified.
	file, err := os.OpenFile("a_copy.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,os.ModePerm)
	if err != nil {
		fmt.Println("文件打开失败：",err)
		return
	}

	//关闭文件
	defer file.Close()

	//写出数据
	by := []byte{65,66,67,78,89}
	n,err:=file.Write(by)
	fmt.Println(n)
	fmt.Println(string(by))


	n,err = file.WriteString("北京欢迎你\n")
	if err != nil {
		fmt.Println("文件写入失败：",err)
		return
	}
	fmt.Println(n)
}
