package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//作为一个更完整的例子，处理函数可以报告它接收到的消息头和表单数据
func handler3(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "%s %s %s\n", request.Method, request.URL, request.Proto)
	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header[%q]=%q\n", k, v)
	}
	fmt.Fprintf(writer, "Host = %q\n", request.Host)
	fmt.Fprintf(writer, "RemotAddr = %q\n", request.RemoteAddr)
	if err := request.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range request.Form {
		fmt.Fprintf(writer, "Form[%q]=%q\n", k, v)
	}
}

//go run WebDemo2.go
//浏览器访问http://localhost:8000/
//GET / HTTP/1.1
//Header["Connection"]=["keep-alive"]
//Header["User-Agent"]=["Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"]
//Header["Accept-Language"]=["zh-CN,zh;q=0.9"]
//Header["Upgrade-Insecure-Requests"]=["1"]
//Header["Sec-Fetch-User"]=["?1"]
//Header["Accept"]=["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3"]
//Header["Sec-Fetch-Site"]=["none"]
//Header["Sec-Fetch-Mode"]=["navigate"]
//Header["Accept-Encoding"]=["gzip, deflate, br"]
//Header["Cookie"]=["_ga=GA1.1.631912846.1571061201"]
//Host = "localhost:8000"
//RemotAddr = "127.0.0.1:49925"
