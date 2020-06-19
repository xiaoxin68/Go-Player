package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) //回调请求调用处理程序
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//处理程序回显请求URL r的路径部分
func handler(writer http.ResponseWriter, request *http.Request) {
	//%q:带引号字符串
	fmt.Fprintf(writer, "URL.Path=%q\n", request.URL.Path)
}

//go run WebDemo.go
//浏览器访问http://localhost:8000、http://localhost:8000/hh
