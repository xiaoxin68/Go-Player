package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64 = 2.0, 4.0, 2.0
	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		x2 := (-b - math.Sqrt(m)) / (2 * a)
		fmt.Printf("有两个解，x1=%f,x2=%f\n", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / (2 * a)
		fmt.Printf("有一个解:%f\n", x1)
	} else {
		fmt.Println("无解")
	}
}
