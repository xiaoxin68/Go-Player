package main

import "fmt"

// 用于测试值传递效果的结构体
type MyData struct {
	complax []int // 测试切片在参数传递中的效果
	
	instance InnerData // 实例分配的innerData
	
	ptr *InnerData // 将ptr声明为InnerData的指针类型
}

// 代表各种结构体字段
type InnerData struct {
	a int
}

// 值传递测试函数
func passByValue(inFunc MyData) MyData {
	
	// 输出参数的成员情况
	fmt.Printf("inFunc value: %+v\n", inFunc)
	
	// 打印inFunc的指针
	fmt.Printf("inFunc ptr: %p\n", &inFunc)
	
	return inFunc
}

func main() {
	
	// 准备传入函数的结构
	in := MyData{
		complax: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}
	
	// 输入结构的成员情况
	fmt.Printf("in value: %+v\n", in)
	
	// 输入结构的指针地址
	fmt.Printf("in ptr: %p\n", &in)
	
	// 传入结构体，返回同类型的结构体
	out := passByValue(in)
	
	// 输出结构的成员情况
	fmt.Printf("out value: %+v\n", out)
	
	// 输出结构的指针地址
	fmt.Printf("out ptr: %p", &out)
}

/*所有的 Data 结构的指针地址都发生了变化，意味着所有的结构都是一块新的内存，
无论是将 Data 结构传入函数内部，还是通过函数返回值传回 Data 都会发生复制行为。
所有的 Data 结构中的成员值都没有发生变化，原样传递，意味着所有参数都是值传递。
Data 结构的 ptr 成员在传递过程中保持一致，表示指针在函数参数值传递中传递的只是指针值，不会复制指针指向的部分。*/