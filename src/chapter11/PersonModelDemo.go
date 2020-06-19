package main

import (
	"./models"
	"fmt"
)

func main() {
	var p = models.NewPerson("zhang")
	p.SetAge(45)
	p.SetSalary(7892.89)
	fmt.Println(*p)
	fmt.Printf("name=%v,age=%v,salary=%v\n",p.Name,p.GetAge(),p.GetSalary())
}
