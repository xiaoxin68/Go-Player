package main

import (
	"fmt"
	"io/ioutil"
)

//写文件（使用ioutil一次将整个文件写出）
func main() {
	 fileName := "C:/Users/zxx/Desktop/b.txt"
	
	 var str = []byte("hello,北京\n")
	
	err := ioutil.WriteFile(fileName,str,0666)
	if err != nil {
		fmt.Println("write file err=%v", err)
	}
}
