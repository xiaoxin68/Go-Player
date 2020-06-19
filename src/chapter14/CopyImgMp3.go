package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//func Copy(dst Writer, src Reader) (written int64, err error)
//将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误。

func CopyFlie(distFileName string,srcFilename string)(written int64, err error)  {
	srcFile,err := os.Open(srcFilename)
	if err != nil{
		fmt.Printf("open flie err=%v\n",err)
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	
	distFile,err := os.OpenFile(distFileName,os.O_WRONLY|os.O_CREATE,0666)
	if err != nil{
		fmt.Printf("open flie err=%v\n",err)
		return
	}
	writer := bufio.NewWriter(distFile)
	defer distFile.Close()
	
	return io.Copy(writer,reader)
}
func main() {
	distFileName,srcFilename := "C:/Users/zxx/Desktop/1.CHM","C:/Users/zxx/Desktop/JDK_API_1_6_zh_CN.CHM"
	written, _ := CopyFlie(distFileName, srcFilename)
	fmt.Println(written)
}
