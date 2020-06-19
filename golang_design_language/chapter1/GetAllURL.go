package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //启动一个协程
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //从ch通道接收
	}
	fmt.Printf("%.2fs elasped\n", time.Since(start).Seconds()) //输出程序运行时间
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url) //产生一个http请求，返回结果存在resp
	if err != nil {
		ch <- fmt.Sprint(err) //发送到通道ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body) //copy函数读取相应的内容，ioutil.Discard输出流进行丢弃；返回字节数
	resp.Body.Close()                                 //关闭body数据，避免资源泄露
	if err != nil {
		ch <- fmt.Sprintf("fetch reading %s:%v\n\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

//go run GetAllURL.go https://hao.360.com/ https://www.baidu.com/ https://www.jd.com/
/*0.71s     227 https://www.baidu.com/
0.74s  110503 https://www.jd.com/
0.89s  668420 https://hao.360.com/
0.90s elasped*/
