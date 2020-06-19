package main

import (
	"fmt"
	"math/rand"
	"time"
)

//临界资源
var tickets = 10

func main() {
	go sale("窗口1")
	go sale("窗口2")
	go sale("窗口3")
	go sale("窗口4")

	time.Sleep(time.Second * 5)
}
func sale(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：", tickets)
			tickets--
		} else {
			fmt.Println(name, "售罄，没票了。。。")
			break
		}
	}
}
