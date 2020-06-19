package main

import "fmt"

//删除切片中的元素
func main() {
	//一、删除开头元素的三种方法
	
	//1、删除开头的元素可以直接移动数据指针：
	a := []int{1, 2, 3}
	a = a[1:] // 删除开头1个元素
	//a = a[N:] // 删除开头N个元素
	fmt.Println(a)
	
	//2、也可以不移动数据指针，但是将后面的数据向开头移动，可以用 append 原地完成
	//（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）
	b := []int{1, 2, 3}
	b = append(b[:0], b[1:]...) // 删除开头1个元素
	//a = append(a[:0], a[N:]...) // 删除开头N个元素
	fmt.Println(b)
	
	//3、copy() 函数来删除开头的元素:copy返回被复制的元素数量
	c := []int{1, 2, 3}
	c = c[:copy(c, c[1:])] // 删除开头1个元素
	//c = c[:copy(c, c[N:])] // 删除开头N个元素
	fmt.Println(c)
	
	//二、从中间位置删除切片中的元素
	
	//1、使用append
	d := []int{1, 2, 3}
	d = append(d[:1], d[1+1:]...)
	//d = append(d[:i], d[i+1:]...) // 删除中间1个元素
	//d = append(d[:i], d[i+N:]...) // 删除中间N个元素
	fmt.Println(d)
	
	//2、使用copy
	e := []int{1, 2, 3}
	e = e[:1+copy(e[1:],e[1+1:])]
	//a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	//a = a[:i+copy(a[i:], a[i+N:])] // 删除中间N个元素
	fmt.Println(e)
	
	//三、从尾部删除
	f := []int{1, 2, 3}
	f = f[:len(f)-1]
	//a = a[:len(a)-1] // 删除尾部1个元素
	//a = a[:len(a)-N] // 删除尾部N个元素
	fmt.Println(f)
	
	//四、删除切片指定位置的元素。
	seq := []string{"a", "b", "c", "d", "e"}
	// 指定删除位置
	index := 2
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
}
