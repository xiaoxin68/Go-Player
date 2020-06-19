package main

import "fmt"

//当发生panic的时候，已经被defer的函数会被执行，还未被defer的函数不会执行
//当外围函数的代码发生了运行时恐慌，只有其中所有已经defer的函数全部执行完毕后，该运行恐慌才会真正被扩展到调用处

//recover语句类似于catch语句，捕获异常
func main() {
	fmt.Println("进入main函数。。。")
	defer fmt.Println("defer 3")
	panicTest()
	defer fmt.Println("defer 4")
	fmt.Println("结束main函数。。。")
}

func panicTest() {
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

//什么情况使用错误表达，什么情况用异常表达？
/*以下情况使用异常处理：
1、空指针引用
2、下标越界
3、除数为零
4、不应该出现的分支，比如default
5、输入兵应该引起函数错误
其他场景使用错误处理机制*/






