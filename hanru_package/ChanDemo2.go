package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string, 4)
	go sentData(ch)
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("读完了。。。", ok)
			break
		}
		fmt.Println("\t读取的数据是:", v)
	}
	fmt.Println("main end...")
}

func sentData(ch chan string) {
	for i := 1; i <= 10; i++ {
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("子channel写入第%d个数据\n", i)
	}
	close(ch)
}
