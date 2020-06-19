package main

import (
	"fmt"
	"io/ioutil"
)
//使用ioutil一次将整个文件读入写出
func main() {
	readName := "C:/Users/zxx/Desktop/a.txt"
	writeName := "C:/Users/zxx/Desktop/b.txt"
	content, err := ioutil.ReadFile(readName)
	if err != nil {
		fmt.Println("read file err=%v", err)
		return
	}
	err = ioutil.WriteFile(writeName,content,0666)
	if err != nil {
		fmt.Println("write file err=%v", err)
	}
}
