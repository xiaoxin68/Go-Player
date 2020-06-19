package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//打开文件
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println("文件打开失败：",err)
		return
	}
	//关闭文件
	defer file.Close()
	//读取数据
	n := -1
	bytes := make([]byte,1024,1024)  //相当于自定义缓冲区
	for{
		n,err =file.Read(bytes)
		if n==0 || err == io.EOF {
			//fmt.Println("读到了文件末尾，结束读取操作")
			break
		}
		fmt.Print(string(bytes[:n]))
	}
}
