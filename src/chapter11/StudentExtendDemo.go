package main

import "fmt"

//学生
type Student struct {
	Name string
	Age int
	Score float64
}

func (stu *Student)Testing()  {
	fmt.Println("学生正在考试")
}

func (stu *Student)ShowInfo()  {
	fmt.Printf("姓名：%v，年龄：%v，成绩：%.2f\n",stu.Name,stu.Age,stu.Score)
}
func (stu *Student)SetScore(s float64)  {
	stu.Score = s
}
//小学生
type Pupil struct {
	Student
}
func (stu *Pupil)Testing()  {
	fmt.Println("小学生正在考试")
}
//中学生
type Middle struct {
	Student
}
func (stu *Middle)Testing()  {
	fmt.Println("中学生正在考试")
}
func main() {
	var p = Pupil{Student{
		Name:  "张三",
		Age:   10,
		Score: 0,
	}}
	p.Testing()
	p.SetScore(90.23)
	p.ShowInfo()
	
	//var m Middle
	var m = &Middle{}
	m.Student.Name = "李四"
	m.Student.Age = 16
	m.Testing()
	m.SetScore(89.50)
	m.ShowInfo()
}
