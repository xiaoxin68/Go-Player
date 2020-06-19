package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//go run GetURL.go https://hao.360.com/
//go run GetURL.go 错误的url

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) //产生一个http请求，返回结果存在resp
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body) //读取整个相应结果并存入b中
		resp.Body.Close()                   //关闭body数据，避免资源泄露
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s:%v\n\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
