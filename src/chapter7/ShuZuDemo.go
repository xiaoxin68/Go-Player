package main

import "fmt"

func main() {
	var weight [5]float64
	weight[0] = 12.3
	weight[1] = 34.12
	weight[2] = 23.45
	weight[3] = 12.54
	weight[4] = 34.56
	var totalWeight float64 = 0.0
	for i := 0;i<len(weight);i++ {
		totalWeight += weight[i]
	}
	fmt.Printf("这些鸡的总体重为%.2f，平均体重为%.2f\n",totalWeight,totalWeight/float64(len(weight)))
}
