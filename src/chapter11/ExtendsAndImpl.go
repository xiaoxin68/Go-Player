package main

import "fmt"

type Monkey struct {
	Name string
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, "生来会爬树")
}

type BirdAdle interface {
	Flying()
}
type Fishable interface {
	Swimming()
}
type LittleMonkey struct {
	Monkey //继承
}

func (littleMonkey *LittleMonkey) Flying() {
	fmt.Println(littleMonkey.Name, "学会了飞翔")
}

func (littleMonkey *LittleMonkey) Swimming() {
	fmt.Println(littleMonkey.Name, "学会了游泳")
}
func main() {
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
