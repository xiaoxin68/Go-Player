package main

import "fmt"

//8个协程计算1+2+……+n  【1-20000】

func writeChan(intChan chan int)  {
	for i:=1;i<=2000;i++  {
		intChan<-i
	}
	close(intChan)
}
func readNum(intChan chan int,resChan chan int,exitChan chan bool)  {
	var sum int
	for{
		val,ok := <-intChan
		if !ok {  //intChan关闭了
			break
		}
		//计算累加和
		sum = 0
		for i:=1;i<=val;i++{
			sum+=i
		}
		resChan<-sum
	}
	//此时还不能close
	fmt.Println("有一个resChan协程因为取不到数据，退出")
	exitChan<-true
}
func main() {
	intChan := make(chan int,1000)
	resChan := make(chan int,1000)
	exitChan := make(chan bool,8)
	//开启一个协程写数据
	go writeChan(intChan)
	//开启8个协程读数据
	for i:=1;i<=8 ;i++  {
		go readNum(intChan,resChan,exitChan)
	}
	//主线程
	go func() {
		for i:=1;i<=8 ;i++{
			<-exitChan
		}
		close(resChan)
	}()
	//输出res中的结果
	for {
		v ,ok:=<-resChan
		if !ok {
			break
		}
		fmt.Printf("累加和=%d\n",v)
	}
	fmt.Println("main线程退出")
}
