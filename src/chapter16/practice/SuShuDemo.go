package main

//将统计素数的任务分配给多个（4个）协程去完成

import (
	"fmt"
)

//向intChan放入1-8000个数
func putNum(intChan chan int)  {
	for i:=1;i<=8000 ;i++  {
		intChan<-i
	}
	//关闭intChan
	close(intChan)
}
//从intChan取出数据，并判断是否为素数，如果是，就放入到primeChan
func primeNum(intChan chan int,primeChan chan int,exitChan chan bool) {
	var flag bool
	for{
		num,ok:=<-intChan
		if !ok { //intChan关闭
			break
		}
		flag = true //假设是素数
		//判断是不是素数
		for i:=2;i<num ;i++  {
			if num%i==0 {//说明不是素数
				flag = false
				break
			}
		}
		if flag {
			primeChan<-num
		}
	}
	fmt.Println("有一个primeChan协程因为取不到数据，退出")
	//这里我们还不能关闭primeChan
	//向exitChan写入true
	exitChan<-true
}

func main()  {
	intChan := make(chan int,1000)
	primeChan := make(chan int,2000)
	//标识退出的管道
	exitChan := make(chan bool,4) //4个
	go putNum(intChan) //开启一个协程，向intChan放入1-8000个数
	//开启4个协程，从intChan取出数据，并判断是否为素数
	for i:=1;i<=4 ;i++  {
		go primeNum(intChan,primeChan,exitChan)
	}
	//主线程
	go func() {
		for i:=1;i<=4 ;i++  {
			<-exitChan
		}
		//当我们从exitChan取出了4个结果，就可以放心关闭primeChan
		close(primeChan)
	}()
	//遍历primeChan取出结果
	for{
		res,ok:=<-primeChan
		if !ok { //primeChan关闭
			break
		}
		fmt.Printf("素数=%d\n",res)
	}
	fmt.Println("main线程退出")
}
