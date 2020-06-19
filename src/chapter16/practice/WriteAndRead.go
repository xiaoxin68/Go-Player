package main

//一个协程读数据，一个协程写数据

import (
	"fmt"
)

//writeData
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intChan <- i
		fmt.Println("writeData ", i)
		//time.Sleep(time.Second)
	}
	//写完intChan后关闭管道
	close(intChan)
}
//readData
func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		//注意啊：并不是intChan一close了它就break，
		//它会一直读，直到读完了才break！！！！
		if !ok {//说明intChan管道已经关闭，退出
			break
		}
		fmt.Printf("readData 读到数据=%v\n", v)
		//time.Sleep(time.Second)
	}
	//读取完毕后，主线程可以退出
	//关闭exitChan管道
	exitChan <- true
	close(exitChan)
}
func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)  //如果管道容量还是10，没有读协程，那么会报错【deadlock!】,因为管道容量不够

	//x, ok := <-ch【只需要看第二个 bool 返回值即可，如果返回值是 false 则表示 ch 已经被关闭。】
	for {
		_, ok := <-exitChan
		if !ok { //说明exitChan管道已经关闭：即已经读完了
			break
		}
	}
}
