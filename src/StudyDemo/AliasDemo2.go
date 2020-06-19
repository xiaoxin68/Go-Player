package main

import "time"

//非本地类型不能定义方法
// 定义time.Duration的别名为MyDuration
type MyDuration = time.Duration

// 为MyDuration添加一个函数
/*func (m MyDuration) EasySet(a string) {
	//cannot define new methods on non-local type time.Duration
}
*/

/*编译器提示：不能在一个非本地的类型 time.Duration 上定义新方法，
非本地类型指的就是 time.Duration 不是在 main 包中定义的，而是在 time 包中定义的，与 main 包不在同一个包中，
因此不能为不在一个包中的类型定义方法。

解决这个问题有下面两种方法：
将第 10 行修改为 type MyDuration time.Duration，也就是将 MyDuration 从别名改为类型；
将 MyDuration 的别名定义放在 time 包中。*/
func main() {
}