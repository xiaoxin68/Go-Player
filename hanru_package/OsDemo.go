package main

import (
	"fmt"
	"os"
)

func main() {
	//Stat返回描述文件f的FileInfo类型值
	fileInfo, err := os.Stat("a.txt")

	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("%T\n", fileInfo)
	//文件名
	fmt.Println(fileInfo.Name())
	//文件大小
	fmt.Println(fileInfo.Size())//字节为单位
	//是否为目录
	fmt.Println(fileInfo.IsDir())
	//修改时间
	fmt.Println(fileInfo.ModTime())
	//权限
	fmt.Println(fileInfo.Mode())//分别是owner group others的权限
}
