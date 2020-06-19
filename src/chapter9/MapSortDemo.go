package main

import (
	"fmt"
	"sort"
)

//go中的map默认是无序的，那么如何排序呢
func main() {
	a := make(map[int]int, 10)
	a[10] = 100
	a[1] = 13
	a[78] = 56
	a[8] = 90
	
	//如果按照map的key的顺序进行排序输出
	//1、先将map的key放入到切片中
	//2、对切片进行排序
	//3、遍历切片，然后按照key来输出map的值
	
	var keys []int
	for key, _ := range a {
		keys = append(keys, key)
	}
	
	sort.Ints(keys)
	
	for _, val := range keys {
		fmt.Printf("map[%d]=%d\n", val, a[val])
	}
}
