package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 3 {
			return
		}
		fmt.Println("i=", i)
	}
	fmt.Println("end")
}
