package main

import (
	"fmt"
	"os"
)

func PathExists(path string)(bool,error)  {
	_,err := os.Stat(path)
	if err == nil {  //说明文件或文件夹存在
		return true,nil
	}
	if os.IsNotExist(err){//说明文件或文件夹不存在
		return false,nil
	}
	return false,err //不确定是否存在
}
func main() {
	exist,_ := PathExists("C:/Users/zxx/Desktop/b.txt")
	fmt.Println(exist)
	
	exist,_ = PathExists("C:/Users/zxx/Desktop/M.txt")
	fmt.Println(exist)
}
