package main
var global *int
func f() {
	var x int
	x = 1
	global = &x
}
func g() {
	y := new(int)
	*y = 1
}
func main() {
	f()
	g()
}
//go run -gcflags "-m -l" TaoYiDemo3.go

//# command-line-arguments
//.\TaoYiDemo3.go:4:6: moved to heap: x
//.\TaoYiDemo3.go:9:10: new(int) does not escape

/*上述代码中，函数 f 里的变量 x 必须在堆上分配，因为它在函数退出后依然可以通过包一级的 global 变量找到，
虽然它是在函数内部定义的。用Go语言的术语说，这个局部变量 x 从函数 f 中逃逸了。

相反，当函数 g 返回时，变量 *y 不再被使用，也就是说可以马上被回收的。因此，*y 并没有从函数 g 中逃逸，
编译器可以选择在栈上分配 *y 的存储空间，也可以选择在堆上分配，然后由Go语言的 GC（垃圾回收机制）回收这个变量的内存空间。

在实际的开发中，并不需要刻意的实现变量的逃逸行为，因为逃逸的变量需要额外分配内存，同时对性能的优化可能会产生细微的影响。*/
