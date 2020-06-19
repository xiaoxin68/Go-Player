package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//在go中需要一个随机数种子，否则返回的数总是固定的
	//time.Now().Unix():返回一个从1997 1-1 00:00:00到现在的一个秒数
	rand.Seed(time.Now().Unix()) //设置随机数种子
	fmt.Println("randomNum:",rand.Intn(100)+1)
	
	var count int = 0
	for {
		ram := 1 + rand.Intn(100)
		count++
		if ram == 99 {
			break
		}
	}
	fmt.Printf("生成99一共经历了%d次", count)
}
