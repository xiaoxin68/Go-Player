package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

//只要它实现sort的全部方法，那么就相当于实现了Sort接口
type HeroSlice []Hero

// Len方法返回集合中的元素个数 :Len() int
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less方法报告索引i的元素是否比索引j的元素小 :Less(i, j int) bool
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	//修改成对name排序
	//return hs[i].Name < hs[j].Name
}

// Swap方法交换索引i和j的两个元素 :Swap(i, j int)
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
func main() {
	//int类型的切片排序
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	
	//结构体类型的切片排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hreo := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hreo)
	}
	//排序前
	fmt.Println(heroes)
	//排序
	sort.Sort(heroes)
	//排序后
	fmt.Println(heroes)
}
