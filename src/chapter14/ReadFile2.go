package main

import (
	"fmt"
	"io/ioutil"
)

//读取文件内容并显示在终端（使用ioutil一次将整个文件读入到内存）
func main() {
	fileName := "C:/Users/zxx/Desktop/hexo.txt"
	
	//使用ioutil一次将整个文件读入到内存
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("read file err=%v", err)
	}
	
	//把读取的文件内容显示到终端
	//fmt.Printf("%v", content)         //[]byte
	fmt.Printf("%v", string(content)) //[]byte
	
	//我们没有显式地Open文件，因此也不需要显式的Close文件
	//因为文件的Open和Close被封装在ReadFile函数内部
}
