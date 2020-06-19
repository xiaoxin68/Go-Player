package main

import "fmt"

//除了在切片头部增加元素，还可以在切片尾部增加元素
func main() {
	//在切片尾部增加元素
	//在切片开头添加元素一般都会导致内存的重新分配，而且会导致已有元素全部被复制 1 次，
	//因此，从切片的开头添加元素的性能要比从尾部追加元素的性能差很多。
	var a = []int{1, 2, 3}
	a = append([]int{0}, a...)          // 在开头添加1个元素
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
	fmt.Println(a)                      //[-3 -2 -1 0 1 2 3]
	
	//实现在切片中间插入元素
	//append 函数返回新切片的特性，所以切片也支持链式操作，我们可以将多个 append 操作组合起来，实现在切片中间插入元素：
	var b []int = []int{1, 2, 3}
	b = append(b[:1], append([]int{100}, b[1:]...)...) // 在第i个位置插入x
	fmt.Println(b)                                     //[1 100 2 3]
	
	//每个添加操作中的第二个 append 调用都会创建一个临时切片，
	// 并将 a[i:] 的内容复制到新创建的切片中，然后将临时创建的切片再追加到 a[:i] 中
	c := []int{1, 2, 3}
	c = append(c[:1], append([]int{-1, -2, -3}, c[1:]...)...) // 在第i个位置插入切片
	fmt.Println(c)                                            //[1 -1 -2 -3 2 3]
}