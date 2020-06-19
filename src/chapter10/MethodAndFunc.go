package main

import "fmt"

//对于函数，如果接收参数值指针，则只能传地址;反之亦然
type FUNC struct {
	Name string
}

func func1(f FUNC)  {
	fmt.Println(f.Name)
}

func func2(f *FUNC)  {
	fmt.Println((*f).Name)
}


//对于方法，如果接收参数值指针，即可传值类型，也可传地址
type METHOD struct {
	Name string
}

func (m METHOD)method1() {
	fmt.Println(m.Name)
}

func (m *METHOD) method2() {
	fmt.Println(m.Name)
	//fmt.Println((*m).Name)  //这两种写法都可以
}
func main() {
	f1 := FUNC{"张三"}
	func1(f1)
	func2(&f1)
	
	m1 := METHOD{"李四"}
	m1.method1()
	(&m1).method1()  //虽然传的是地址，但本质上还是值拷贝！！！
	m1.method2()  //虽然传的是不是地址，但本质上仍然是地址拷贝！！！
	(&m1).method2()
}
