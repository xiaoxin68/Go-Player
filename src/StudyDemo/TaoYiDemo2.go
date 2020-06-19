package main

import "fmt"

// 声明空结构体测试结构体逃逸情况
type Data struct {
}

func dummy2() *Data {
	// 实例化c为Data类型
	var c Data
	
	//返回函数局部变量地址
	return &c
}

func main() {
	fmt.Println(dummy2())
}

//go run -gcflags "-m -l" main.go

//# command-line-arguments
//.\TaoYiDemo2.go:11:6: moved to heap: c
//.\TaoYiDemo2.go:18:13: ... argument does not escape

/*moved to heap: c:将 c 移到堆中。
这句话表示，Go 编译器已经确认如果将变量 c 分配在栈上是无法保证程序最终结果的，
如果这样做，dummy() 函数的返回值将是一个不可预知的内存地址，
这种情况一般是 C/C++ 语言中容易犯错的地方，引用了一个函数局部变量的地址。*/

/*在使用Go语言进行编程时，Go语言的设计者不希望开发者将精力放在内存应该分配在栈还是堆的问题上，
编译器会自动帮助开发者完成这个纠结的选择，但变量逃逸分析也是需要了解的一个编译器技术，
这个技术不仅用于Go语言，在 Java 等语言的编译器优化上也使用了类似的技术。

编译器觉得变量应该分配在堆和栈上的原则是：
变量是否被取地址；
变量是否发生逃逸。*/