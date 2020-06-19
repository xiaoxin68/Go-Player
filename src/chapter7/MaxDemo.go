package main

import "fmt"

func main() {
	var num = [...]int{1,23,4,56,7,45,345,34,3,242}
	var max,index int = num[0],0
	for i:=0;i<len(num) ;i++  {
		if num[i] > max{
			max = num[i]
			index = i
		}
	}
	fmt.Println("max=",max,",index=",index)
}

