package main

import (
	"net/http"
)

func main() {
	//使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录。
	http.Handle("/", http.FileServer(http.Dir(".")))
	//默认的 HTTP 服务侦听在本机 8080 端口
	http.ListenAndServe(":8081", nil)
}

//在浏览器里输入http://127.0.0.1:8081即可浏览文件，这些文件正是当前目录在HTTP服务器上的映射目录。
