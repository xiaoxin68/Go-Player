package main

import (
	"errors"
	"fmt"
)

func main() {
	test02()
	fmt.Println("main go on")
}

func readConf(name string)(err error)  {
	if name == "config.ini"{
		//进行读取文件后的操作
		return nil
	}else {
		return errors.New("文件读取错误")
	}
}
func test02()  {
	err := readConf("hello.ini")
	if err != nil{
		//如果读取文件错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test02 go on")
}
