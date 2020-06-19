package main

import "fmt"

func main()  {
	num := 30
	if num >= 20 {
		goto here
	}
	fmt.Println("there")
	here:
		fmt.Println("here")
	fmt.Println("end")
}
