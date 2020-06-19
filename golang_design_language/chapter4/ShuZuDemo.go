package main

import "fmt"

func main() {
	b :=[32]byte{'s','2','d'}
	zero1(&b)
	fmt.Println(b)
}
func zero1(ptr *[32]byte)  {
	for i:=range ptr{
		ptr[i] = 0
	}
}

func zero2(ptr *[32]byte)  {
	*ptr = [32]byte{}
}