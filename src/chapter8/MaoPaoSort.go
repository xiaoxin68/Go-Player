package main

import "fmt"

func main() {
	numArr := []int{12, 34, 56, 7, 67, 36}
	fmt.Println(numArr)
	maopao(numArr)
	fmt.Println(numArr)
	
	numArr2 := [5]int{3,5,1,3,8}
	fmt.Println(numArr2)
	maopao2(&numArr2)
	fmt.Println(numArr2)
}
//切片的方式
func maopao(numArr []int) {
	for i := 0; i < len(numArr)-1; i++ {
		for j := 0; j < len(numArr)-i-1; j++ {
			if numArr[j] > numArr[j+1] {
				temp := numArr[j]
				numArr[j] = numArr[j+1]
				numArr[j+1] = temp
			}
		}
	}
}

//数组的方式
func maopao2(numArr *[5]int) {
	for i := 0; i < len(*numArr)-1; i++ {
		for j := 0; j < len(*numArr)-i-1; j++ {
			if (*numArr)[j] > (*numArr)[j+1] {
				temp := (*numArr)[j]
				(*numArr)[j] = (*numArr)[j+1]
				(*numArr)[j+1] = temp
			}
		}
	}
}

