package main

import "fmt"

func main() {
	//var numArr = [...][3]int{{1,2,3},{4,5,6},{7,8,9}}
	var numArr = [...][4]int{{1,2,3,4},{5,6,7,8},{8,10,11,12},{13,14,15,16}}
	fmt.Println(numArr)
	for i := 0; i < len(numArr);i++  {
		for j:=0;j<i ;j++  {
			temp := numArr[i][j]
			numArr[i][j] = numArr[j][i]
			numArr[j][i] = temp
		}
	}
	fmt.Println(numArr)
}
