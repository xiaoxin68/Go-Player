package main

import "fmt"

func main() {
	var num = [...]int{12,23,4,56,7,45,345,34,3,242}
	sum := 0
	avg := 0.0
	for _,val := range num {
		sum+=val
	}
	avg =float64(sum)/ float64(len(num))
	fmt.Printf("和为%d，平均值为%.2f",sum,avg)
}
