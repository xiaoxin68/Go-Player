package main

import "fmt"

func main() {
	var slice []string = []string{"黑包子", "宋江", "李逵"}
	//1、for遍历
	for i:=0;i<len(slice) ;i++  {
		fmt.Printf("%s ",slice[i])
	}
	fmt.Println()
	//2、for-range遍历
	for _,val := range slice{
		fmt.Printf("%s ",val)
	}
	fmt.Println()
}
