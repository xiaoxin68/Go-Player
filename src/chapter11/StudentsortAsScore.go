package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*type Interface interface {
	// Len方法返回集合中的元素个数
	Len() int
	// Less方法报告索引i的元素是否比索引j的元素小
	Less(i, j int) bool
	// Swap方法交换索引i和j的两个元素
	Swap(i, j int)
}*/
type MyStudent struct {
	Name  string
	Age   int
	Score float64
}
type StudentSlice []MyStudent

func (stus StudentSlice) Len() int {
	return len(stus)
}
func (stus StudentSlice) Less(i, j int) bool {
	return stus[i].Score < stus[j].Score
}
func (stus StudentSlice) Swap(i, j int) {
	stus[i], stus[j] = stus[j], stus[i]
}
func main() {
	var stus StudentSlice
	for i := 0; i <= 10; i++ {
		stu := MyStudent{
			Name:  fmt.Sprintf("Name%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: float64(rand.Intn(100)),
		}
		stus = append(stus, stu)
	}
	fmt.Println(stus)
	sort.Sort(stus)
	fmt.Println(stus)
}
