package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
//使用带缓冲区的方式
func main() {
	readFile,err := os.Open("C:/Users/zxx/Desktop/a.txt")
	writeFile,err2 := os.OpenFile("C:/Users/zxx/Desktop/b.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil || err2 != nil{
		fmt.Println("open file err=", err)
		fmt.Println("open file err=", err2)
		return
	}
	//关闭文件
	defer readFile.Close()
	defer writeFile.Close()
	//缓冲区大小
	const defaultBuffer  = 4096
	//创建一个*Reader，带缓冲的
	reader := bufio.NewReader(readFile)
	writer := bufio.NewWriter(writeFile)
	//循环读取文件内容
	for{
		str,err := reader.ReadString('\n')  //读到一个换行就结束
		if err == io.EOF {//io.EOF表示文件的末尾
			break
		}
		//写入文件
		writer.WriteString(str)
	}
	writer.Flush()
}
