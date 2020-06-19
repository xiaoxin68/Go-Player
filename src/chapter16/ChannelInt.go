package main

import "fmt"

func main() {
	var intChan chan int = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	intChan <- 30
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Printf("num1=%d,num2=%d,num3=%d\n",num1,num2,num3) //num1=10,num2=20,num3=30
}
