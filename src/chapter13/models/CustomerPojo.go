package models

import "fmt"

type Customer struct {
	id     int
	name   string
	gender string
	age    int
	phone  string
	email  string
}
//返回*Customer的方法
func NewCustomer(id int, name string, gender string, age int, phone string, email string) (*Customer) {
	return &Customer{
		id:     id,
		name:   name,
		gender: gender,
		age:    age,
		phone:  phone,
		email:  email,
	}
}
//没有id的方法
func NewCustomer2(name string, gender string, age int, phone string, email string) (*Customer) {
	return &Customer{
		name:   name,
		gender: gender,
		age:    age,
		phone:  phone,
		email:  email,
	}
}
//setid
func (this *Customer)SetId(id int)  {
	this.id = id
}
func (this *Customer)GetId() int {
	return this.id
}
//返回顾客信息的方法
func (this *Customer) GetInfo ()string{
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.id,this.name,this.gender,this.age,
		this.phone,this.email)
	return info
}
