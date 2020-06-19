package main

import "fmt"

func main() {
	//三个班五名学生的平均成绩
	var scores [3][5]float64
	for i:=0;i<len(scores) ;i++  {
		for j:=0;j<len(scores[i]);j++ {
			fmt.Printf("请输入第%d个班第%d名学生的成绩\n",i+1,j+1)
			fmt.Scanln(&scores[i][j])
		}
	}
	
	var total float64 = 0.0
	for i,v:=range scores {
		temp := 0.0
		for _,val := range v {
			temp+=val
		}
		total+=temp
		fmt.Printf("第%d个班学生的平均成绩是%.2f\n",i+1,temp/float64(len(v)))
	}
	fmt.Printf("全部学生的平均成绩是%.2f\n",total/float64(len(scores)*len(scores[0])))
}
