package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[int]string)
	m[1] = "红孩儿"
	m[2] = "小站费"
	m[3] = "英语课"
	m[4] = "信息课"
	m[5] = "lalalla"
	m[6] = "q4324"


	fmt.Println("===========原始输出===========")
	for i, v := range m {
		fmt.Printf("index:%d,val:%s\n", i, v) //输出的结果是无序的
	}

	fmt.Println("===========key排序后输出===========")

	keys := make([]int, 0, len(m))
	for i, _ := range m {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for _, val := range keys {
		fmt.Printf("index:%d,val:%s\n", val, m[val])
	}
}
