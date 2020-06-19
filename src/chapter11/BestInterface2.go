package main

import "fmt"

type S struct {
	Name string
	Age int
	Score float64
}

func TypeJudge(items ...interface{}) {
	for index, val := range items {
		switch val.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index+1, val)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", index+1, val)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", index+1, val)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n", index+1, val)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index+1, val)
		case S:
			fmt.Printf("第%v个参数是S类型，值是%v\n", index+1, val)
		case *S:
			fmt.Printf("第%v个参数是*S类型，值是%v\n", index+1, *(val.(*S)))
		default:
			fmt.Printf("第%v个参数是类型不确定，值是%v\n", index+1, val)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	var ch byte = 'a'
	n4 := 300
	stu := S{"张三",20,90.25}
	stu2 := &stu
	TypeJudge(n1, n2, n3, name, address, ch, n4,stu,stu2)
}
