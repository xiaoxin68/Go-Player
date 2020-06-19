package main

import "fmt"

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := a[:5]  //[1,2,3,4,5]
	s2 := a[3:7] //[4,5,6,7]
	s3 := a[:]
	fmt.Println("==================原来的数组和切片====================")
	fmt.Println(a)  //[1 2 3 4 5 6 7 8 9 10]
	fmt.Println(s1) //[1 2 3 4 5]
	fmt.Println(s2) //[4 5 6 7]
	fmt.Println(s3) //[1 2 3 4 5 6 7 8 9 10]

	fmt.Println("==================切片和容量====================")
	fmt.Printf("len=%d,cap=%d\n", len(a), cap(a))   //len=10,cap=10
	fmt.Printf("len=%d,cap=%d\n", len(s1), cap(s1)) //len=5,cap=10
	fmt.Printf("len=%d,cap=%d\n", len(s2), cap(s2)) //len=4,cap=7
	fmt.Printf("len=%d,cap=%d\n", len(s3), cap(s3)) //len=10,cap=10

	fmt.Println("==================修改数组中的元素====================")
	a[3] = 100
	fmt.Println(a)  //[1 2 3 100 5 6 7 8 9 10]
	fmt.Println(s1) //[1 2 3 100 5]
	fmt.Println(s2) //[100 5 6 7]
	fmt.Println(s3) //[1 2 3 100 5 6 7 8 9 10]

	fmt.Println("==================修改切片中的元素====================")
	s1[4] = 200
	fmt.Println(a)  //[1 2 3 100 200 6 7 8 9 10]
	fmt.Println(s1) //[1 2 3 100 200]
	fmt.Println(s2) //[100 200 6 7]
	fmt.Println(s3) //[1 2 3 100 200 6 7 8 9 10]

	fmt.Println("==================向切片中appen元素（未超过cap）====================")
	s1 = append(s1, 2, 2, 2)
	fmt.Println(a)  //[1 2 3 100 200 2 2 2 9 10]  //数组仍旧改变
	fmt.Println(s1) //[1 2 3 100 200 2 2 2]
	fmt.Println(s2) //[100 200 2 2]
	fmt.Println(s3) //[1 2 3 100 200 2 2 2 9 10]

	//由于向切片中append元素超过了切片的容量，会扩容，扩容后的s1指向一个新的内存空间，所以对s1的修改不会影响数组和其他切片了~~~反之亦然
	fmt.Println("==================向切片中appen元素（超过cap）====================")
	s1 = append(s1, 3, 3, 3, 3, 3)
	fmt.Println(a)  //[1 2 3 100 200 2 2 2 9 10]
	fmt.Println(s1) //[1 2 3 100 200 2 2 2 3 3 3 3 3]
	fmt.Println(s2) //[100 200 2 2]
	fmt.Println(s3) //[1 2 3 100 200 2 2 2 9 10]
}
