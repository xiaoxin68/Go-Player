package main

import "fmt"

type MethodUtil struct {
}

func (mm MethodUtil) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func (mm *MethodUtil) Print2(m,n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func main() {
	var mm = MethodUtil{}
	mm.Print()
	(&mm).Print2(3,5)
}
