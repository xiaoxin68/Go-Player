package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机生成五个数，数组反转
func main() {
	rand.Seed(time.Now().UnixNano())
	
	//随机生成五个数
	var nums [6]int
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(100) //(0,100)
	}
	
	fmt.Println(nums)
	
	i := 0
	j := len(nums) - 1
	for {
		if i < len(nums)/2 {
			//交换
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			i++
			j--
		}else {
			break
		}
		
	}
	fmt.Println(nums)
}
