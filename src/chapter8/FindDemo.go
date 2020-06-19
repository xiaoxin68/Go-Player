package main

import "fmt"

func main() {
	arr := []string{"张三","李四","王二","马五"}
	fmt.Println(shunxuFind(arr,"王二"))
	fmt.Println(shunxuFind(arr,"小新"))
	
	numArr := []int{2,4,5,7,8,9}
	fmt.Println(dividFind(numArr,7))
	fmt.Println(dividFind(numArr,1))
}
//顺序查找
func shunxuFind(arr []string,str string)int  {
	res := -1
	for index,val := range arr{
		if val == str{
			res = index
			break
		}
	}
	return res
}

//二分查找
func dividFind(arr []int,num int)int  {
	left,right := 0, len(arr)
	for left <= right{
		mid := left + (right-left)/2
		if arr[mid] == num {
			return mid
		}else if arr[mid] > num {
			right = mid-1
		}else {
			left = mid+1
		}
	}
	return -1
}