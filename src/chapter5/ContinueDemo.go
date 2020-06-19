package main

import "fmt"

func main() {
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			if j == 2 {
				continue
			}
			fmt.Println("j=", j)
		}
	}
	
	fmt.Println("==================")
	
	//可通过标签指明要终止到哪一层循环
label2:
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			if j == 2 {
				continue label2
			}
			fmt.Println("j=", j)
		}
	}
}
