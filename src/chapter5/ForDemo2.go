package main

import "fmt"

func main()  {
	var sum,count int = 0,0
	for i:=1;i<=100 ;i++  {
		if i%9 == 0{
			sum += i
			count++
		}
	}
	fmt.Printf("1-100是9的倍数的数一共%d个，和为%d",count,sum)
}
