package main

import "fmt"

func main() {

	//注意到：数组是值传递；
	//数组的地址和首元素的地址是一致的，即数组的地址指向首个元素的地址
	//数组进入函数后，数组的地址和首元素的地址都会重新分配
	fmt.Println("================数组================")
	arr := [4]int{1, 2, 3, 4}
	fmt.Printf("mian里arr=%v，arr地址=%p，arr[0]地址=%v\n", arr, &arr, &arr[0])
	arrChange(arr)
	fmt.Printf("mian里arr=%v，arr地址=%p，arr[0]地址=%v\n", arr, &arr, &arr[0])

	//注意到：切片是引用传递；
	//切片的地址和首元素的地址是不是一致的【切片本身有地址，首元素也有地址】
	//切片进入函数后，切片的地址会虫子呢分配，但是首元素的地址始终保持一致；这就是切片可以改变的原因
	fmt.Println("================切片================")
	slice := []int{1, 2, 3}
	fmt.Printf("mian里slice=%v，slice地址=%p，slice[0]地址=%v,cap=%d\n", slice, &slice, &slice[0], cap(slice))
	sliceChange(slice)
	fmt.Printf("mian里slice=%v，slice地址=%p，slice[0]地址=%v,cap=%d\n", slice, &slice, &slice[0], cap(slice))
}
func arrChange(arr [4]int) {
	fmt.Printf("函数里arr=%v，arr地址=%p，arr[0]地址=%v\n", arr, &arr, &arr[0])
	arr[0] = 10
	fmt.Printf("函数里arr=%v，arr地址=%p，arr[0]地址=%v\n", arr, &arr, &arr[0])
}

//这里也就是说，虽然slice是引用传递，但是slice会涉及到扩容机制，一旦扩容了，切片将指向新的内存单元，
//也就是说，此时对形参的修改将无法改变实参了，所以应当避免这种情况
func sliceChange(slice []int) {
	fmt.Printf("函数里slice=%v，slice地址=%p，slice[0]地址=%v,cap=%d\n", slice, &slice, &slice[0], cap(slice))
	slice[0] = 10 //在扩容之前修改的，有实参效
	slice = append(slice, 20, 30)
	slice[0] = 100 //在扩容之后修改的，对实参无效
	fmt.Printf("函数里里slice=%v，slice地址=%p，slice[0]地址=%v,cap=%d\n", slice, &slice, &slice[0], cap(slice))
}
