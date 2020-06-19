package models

import "fmt"

type person struct {
	Name  string
	age   int
	slary float64
}

func NewPerson(n string) *person { //类似于构造函数
	return &person{
		Name: n,
	}
}

func (p *person) SetAge(a int) { //类似于set方法
	if a > 0 && a < 150 {
		p.age = a
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person) GetAge() int { //类似于get方法
	return p.age
}

func (p *person) SetSalary(s float64) { //类似于set方法
	if s >= 3000 && s <= 30000 {
		p.slary = s
	} else {
		fmt.Println("薪水范围不正确")
	}
}

func (p *person) GetSalary() float64 { //类似于get方法
	return p.slary
}

