package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d",x)
	z := strconv.FormatInt(int64(x),10)  //Itoa
	fmt.Println(y)
	fmt.Println(z)

	m ,_:= strconv.ParseInt(y,10,64)  //Atoi
	fmt.Println(m)
}
