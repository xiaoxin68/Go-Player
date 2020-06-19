package main

import (
	"fmt"
	"go/types"
)

func main() {
	var by byte
	fmt.Println("请输入一个字母【a-z】：")
	fmt.Scanf("%c", &by)
	
	//这里可以使用常量值、变量值、带返回值的函数都可以
	switch test(by) + 1 {
	case 'a':
		fmt.Println("周一")
	case 'b', 'c', 'd':
		fmt.Println("周二")
	case 'e':
		fmt.Println("周三")
	default:
		fmt.Println("周四")
	}
	
	//switch后面也可以不带内容，功能就和if语句一样了，
	// 也可以在switch后面定义变量【不推荐
	score := 87
	switch {
	case score > 90:
		fmt.Println("优秀")
	case score >= 70 && score <= 90:
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}
	
	//switch穿透：结果如下
	//优秀~
	//良好~
	//一般~
	switch grade := 92;{
	case grade >= 90:
		fmt.Println("优秀~")
		fallthrough   //默认穿透一层
	case grade >=80 && grade < 90:
		fmt.Println("良好~")
		fallthrough
	case grade >= 60 && grade <80:
		fmt.Println("一般~")
	default:
		fmt.Println("不及格~")
	}
	
	//用于type-switch来判断某法interface变量实际指向的变量类型
	var  x  interface{}
	var y = 10.0
	x = y
	switch i:=x.(type) {
	case types.Nil:
		fmt.Printf("x的类型为：%T",i)
	case int:
		fmt.Println("x的类型为int型")
	case float64:
		fmt.Println("x的类型为float64型")
	case func(int) float64:
		fmt.Println("x的类型为func(int)型")
	case bool,string:
		fmt.Println("x的类型为bool或string型")
	default:
		fmt.Println("未知型")
	}
}

func test(char byte) byte {
	return char + 1
}
