package main

import (
	"fmt"
	"net"
)

//根据错误获取判断错误类型的方法
func main() {
	host, err := net.LookupHost("www.dgjashsda.com")
	if err != nil {
		fmt.Println(err)
		if ins, ok := err.(*net.DNSError); ok {
			if ins.Timeout() {
				fmt.Println("操作超时")
			} else if ins.Temporary() {
				fmt.Println("临时性错误")
			} else {
				fmt.Println("通常错误")
			}
		}
		return
	}
	fmt.Println("成功", host)
}
