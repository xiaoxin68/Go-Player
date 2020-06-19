package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//读取文件内容并显示在终端（使用带缓冲区的方式）
func main() {
	file,err := os.Open("C:/Users/zxx/Desktop/hexo.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	//关闭文件
	defer file.Close()
	
	//缓冲区大小
	const defaultBuffer  = 4096
	
	//创建一个*Reader，带缓冲的
	reader := bufio.NewReader(file)
	
	//循环读取文件内容
	for{
		str,err := reader.ReadString('\n')  //读到一个换行就结束
		if err == io.EOF {//io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}
	fmt.Println("文件读取结束了。。。")
}
