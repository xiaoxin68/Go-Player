package main

import "fmt"

//当发生panic的时候，已经被defer的函数会被执行，还未被defer的函数不会执行
//recover语句类似于catch语句，捕获异常
//看对比PanicDemo.go查看执行结果的差异
func main() {
	//位置一
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序恢复啦。。。")
		}
	}()

	fmt.Println("进入main函数。。。")
	defer fmt.Println("defer 3")
	RecoverTest()
	defer fmt.Println("defer 4")
	fmt.Println("结束main函数。。。")
}

//在defer的函数中，执行recover调用会取回传至panic调用的错误值，恢复正常执行，停止恐慌过程。
//若recover在defer的函数之外被调用，它将不会停止恐慌过程序列。【所以说：一般recover就是在defer中使用，回复panic】
func RecoverTest() {
	/*defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "程序恢复啦。。。")
		}
	}()*/
	fmt.Println("进入函数中。。。。")
	defer fmt.Println("defer 1")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			panic("panic错误")
		}
		fmt.Println(i)
	}
	fmt.Println()
	defer fmt.Println("defer 2")
}

//没有recover的输出结果：当外围函数的代码发生了panic，只有其中所有已经defer的函数全部执行完毕后，该panic才会真正被扩展到调用处
/*进入main函数。。。
进入函数中。。。。
1
2
3
4
defer 1
defer 3
panic: panic错误*/

//有recover的输出结果：当外围函数的代码发生了panic，按照先进后出的原则执行已经defer的函数，如果遇到recover
//放在位置一：
/*进入main函数。。。
进入函数中。。。。
1
2
3
4
defer 1
defer 3
panic错误 程序恢复啦。。。*/

//放在位置二：从恢复点继续往下执行
/*进入main函数。。。
进入函数中。。。。
1
2
3
4
defer 1
panic错误 程序恢复啦。。。
结束main函数。。。
defer 4
defer 3*/
