package utils

import "fmt"

var Age int = initAge()
var Name string = initName()

func init()  {
	fmt.Println("util-init")
	Age = 18
	Name = "星星"
}

func initAge() int {
	fmt.Println("age-init")
	return 10
}
func initName() string {
	fmt.Println("name-init")
	return "lalala"
}
