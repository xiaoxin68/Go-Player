package main

import "fmt"

func main() {
	//创建一个可以存放3个int类型的管道
	var intChannel chan int
	intChannel = make(chan int, 3)
	
	fmt.Printf("intChannel的值=%v intChannel本身的地址=%p\n", intChannel, &intChannel) //intChannel的值=0xc00007a080 intChannel本身的地址=0xc000006028
	
	//向管道写数据
	intChannel <- 10
	num := 211
	intChannel <- num
	intChannel <- 50
	//注意：当我们给管道写入数据时，不能超过其容量
	
	//看看管道的长度和容量
	fmt.Printf("channel len=%v cap=%v\n", len(intChannel), cap(intChannel)) //channel len=3 cap=3
	
	//从管道中读取数据
	var num2 int
	num2 = <-intChannel
	fmt.Println("num2=", num2)                                              //num2= 10
	fmt.Printf("channel len=%v cap=%v\n", len(intChannel), cap(intChannel)) //channel len=2 cap=3
	
	//在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报fatal error: all goroutines are asleep - deadlock!
	num3 := <-intChannel
	num4 := <-intChannel
	//num5 := <-intChannel
	fmt.Println("num3=", num3, "num4=", num4 /*, "num3=", num5*/) //num3= 211 num4= 50
	
}
