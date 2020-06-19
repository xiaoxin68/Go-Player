package main

import "fmt"

func main() {
	var arr  [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	
	//二维数组遍历方式一
	for i:=0;i<len(arr);i++ {
		for j:=0;j< len(arr[i]);j++  {
			fmt.Printf("%d ",arr[i][j])
		}
		fmt.Println()
	}
	
	//二维数组遍历方式二
	for i1,v1:=range arr{
		for i2,v2 := range v1 {
			fmt.Printf("index(%d,%d)=%d ",i1,i2,v2)
		}
		fmt.Println()
	}
}
