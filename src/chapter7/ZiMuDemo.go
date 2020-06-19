package main

import "fmt"

func main() {
	var myChars [26]byte
	for i:=0;i<len(myChars);i++ {
		myChars[i] = 'A'+byte(i)
	}
	
	for i:=0;i<len(myChars);i++ {
		fmt.Printf("%c ",myChars[i])
	}
	fmt.Println()
}
