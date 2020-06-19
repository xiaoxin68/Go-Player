package main

import "fmt"

func main() {
	var totlaSum float64 = 0.0
	var banji, renshu int = 3, 5
	for i := 1; i <= banji; i++ {
		sum := 0.0
		for j := 1; j <= renshu; j++ {
			var score float64
			fmt.Printf("第%d个班第%d名学生的成绩是：", i, j)
			fmt.Scanln(&score)
			sum += score
		}
		totlaSum += sum
		fmt.Printf("第%d个班学生的平均成绩是：%f\n",i,sum/float64(renshu))
	}
	fmt.Printf("这%d个班学生的平均成绩是：%f\n", banji, totlaSum/float64(renshu*banji))
}
