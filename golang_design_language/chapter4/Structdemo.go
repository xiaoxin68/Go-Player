package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID int
	Name string
	Address string
	DoB time.Time
	Position string
	Salary int
	ManagerId int
}
func main() {
	var delibert Employee
	delibert.Salary -= 5000

	position := &delibert.Position
	*position += "Senior"

	fmt.Println(delibert)

	var employeeTheMonth *Employee = &delibert
	employeeTheMonth.Position += "(proactive)"

	fmt.Println(delibert)
	fmt.Println(employeeTheMonth)

	fmt.Println(EmployeeById(delibert.ManagerId).Position)

	id := delibert.ID
	EmployeeById(id).Salary = 100  //注意：若EmployeeById的返回值变成了Employee而不是*Employee，那么这句话编译不通过。因为赋值表达式的左侧无法识别一个变量

	fmt.Println(delibert)
}
func EmployeeById(id int)  *Employee{
	return &Employee{
		ID: id,
	}
}