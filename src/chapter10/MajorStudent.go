package main

import "fmt"

type MajorStudent struct {
	name string
	gender string
	age int
	id int
	score float64
}

func (stu *MajorStudent)say()string  {
	str := fmt.Sprintf("姓名：%v，性别：%v，年龄：%v，id：%v，成绩：%v\n",stu.name,stu.gender,stu.age,stu.id,stu.score)
	return str
}
func main() {
	var stu = MajorStudent{"小明","男",15,110,90.78}
	fmt.Println(stu.say())
}
