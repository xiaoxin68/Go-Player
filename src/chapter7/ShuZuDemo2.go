package main

import "fmt"

func main() {
	var score [5]float64
	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d位同学的成绩\n", i+1)
		fmt.Scanln(&score[i])
	}
	for i := 0; i < len(score); i++ {
		fmt.Printf("第%d位学生的成绩是%.2f\n", i+1, score[i])
	}
}
