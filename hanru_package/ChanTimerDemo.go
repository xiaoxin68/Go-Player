package main

import (
	"fmt"
	"time"
)

func main() {
	//NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间。
	timer := time.NewTimer(time.Second*3)
	fmt.Printf("%T\n",timer) //*time.Timer
	fmt.Println("now:",time.Now())  //2020-04-30 23:31:26.4722188 +0800 CST m=+0.021988401

	//次数等待chan中的数值，会等待3秒
	ch := timer.C
	fmt.Println(<-ch) //2020-04-30 23:31:29.4723962 +0800 CST m=+3.022165801

	//停止timer的执行：
		//Stop停止Timer的执行。如果停止了t会返回真；如果t已经被停止或者过期了会返回假。
		//Stop不会关闭通道t.C，以避免从该通道的读取不正确的成功。
	timer2 := time.NewTimer(time.Second*5)
	go func() {
		<-timer2.C
		fmt.Println("timer2结束了")
	}()
	time.Sleep(time.Second*3)
	flag:=timer2.Stop()
	if flag {
		fmt.Println("timer2停止了")
	}
}
