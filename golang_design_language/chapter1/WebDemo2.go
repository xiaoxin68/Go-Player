package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int
func main() {

	http.HandleFunc("/", handler2) //回调请求调用处理程序
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//回显目前为止调用的次数
func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count %d\n",count)
	mu.Unlock()
}
//处理程序回显请求的URL路径部分
func handler2(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(writer, "URL.Path=%q\n", request.URL.Path)
}

//go run WebDemo2.go
//浏览器访问http://localhost:8000、http://localhost:8000/hh多次访问
//访问 http://localhost:8000/count 查看次数
