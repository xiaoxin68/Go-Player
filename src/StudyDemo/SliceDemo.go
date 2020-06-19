package main

import "fmt"

//切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……，代码如下：
func main() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}
}

//通过查看代码输出，可以发现一个有意思的规律：切片长度 len 并不等于切片的容量 cap
/*len: 1  cap: 1 pointer: 0xc00000a0c0
len: 2  cap: 2 pointer: 0xc00000a100
len: 3  cap: 4 pointer: 0xc00000e360
len: 4  cap: 4 pointer: 0xc00000e360
len: 5  cap: 8 pointer: 0xc000086000
len: 6  cap: 8 pointer: 0xc000086000
len: 7  cap: 8 pointer: 0xc000086000
len: 8  cap: 8 pointer: 0xc000086000
len: 9  cap: 16 pointer: 0xc00008e000
len: 10  cap: 16 pointer: 0xc00008e000*/
