package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn,err := net.Dial("tcp","192.168.1.103:8888")
	if err !=nil {
		fmt.Println("client dail err=",err)
		return
	}
	
	//功能1：客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin)  //os.Stdin代表标准输入
	
	//在终端输入exit表示退出
	for{
		//从终端读取一行用户输入，并发送给服务器
		line,err := reader.ReadString('\n')
		if err !=nil {
			fmt.Println("ReadString err=",err)
		}
		
		line = strings.Trim(line," \r\n")
		if line == "exit"{
			fmt.Println("客户端退出。。。")
			break
		}
		
		//将line发送给服务器
		n,err := conn.Write([]byte(line+"\n"))
		if err !=nil {
			fmt.Println("conn.Write err=",err)
		}
		fmt.Printf("客户端发送了%d字节的数据\n",n)
	}
}
