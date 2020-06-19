package main

import "fmt"

func main() {
	a := 10
	defer fmt.Println("defer1", a) //10、defer2 20
	a = 20
	defer fmt.Println("defer2", a) //9、defer2 20
	fmt.Println("main1", a)        //1、main1 20

	s := []int{1, 2, 3}
	defer funcCall(s)       //8、funcCall1 [100 2 3]   funcCall2 [100 2 3]
	fmt.Println("main2", s) //2、main2 [1 2 3]

	test(s)
	fmt.Println("main3", s) //7、main3 [100 2 3]
}

func test(a []int) {
	fmt.Println("test1", a) //3、test1 [1 2 3]
	defer funcCall(a)
	fmt.Println("test2", a) //4、test2 [1 2 3]
	a[0] = 200
}

func funcCall(a []int) {
	fmt.Println("funcCall1", a) //5、funcCall1 [200 2 3]
	a[0] = 100
	fmt.Println("funcCall2", a) //6、funcCall2 [100 2 3]
}
