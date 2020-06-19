package main

import (
	"bufio"
	"fmt"
	"os"
)

//写文件（（使用带缓冲区的方式）
func main() {
	filePath := "C:/Users/zxx/Desktop/a.txt"
	// 只写模式打开文件, 如果不存在将创建一个新文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	//如果不要覆盖原来的内容：O_TRUNC
	//追加内容： O_APPEND
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//即时关闭file句柄
	defer file.Close()
	//准备写入5句话
	str := "hello,北京\n"
	//写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 1; i <= 5; i++ {
		writer.WriteString(str)
	}
	//因为writer是带缓存的，因此在调用WriteString方法时，其实内容是先写入到缓存中，所以需要调用Flush方法，将缓存中的数据真正写入到文件中
	//否则文件中会没有数据！！！
	writer.Flush()
}
