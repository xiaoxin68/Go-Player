package main

import "fmt"

func main() {
	var slice []string = []string{"黑包子", "宋江", "李逵"}
	fmt.Println("slice=",slice)  //slice= [黑包子 宋江 李逵]
	
	//append内置函数可以对切片进行动态增长
	slice = append(slice, "及时雨","包青天","李白")
	fmt.Println("slice=",slice) //slice= [黑包子 宋江 李逵 及时雨 包青天 李白]
	
	//切片可以继续切片
	slice2 :=  slice[1:4]
	fmt.Println("slice2=",slice2) //slice2= [宋江 李逵 及时雨]
	
	//通过切片追加的方式，将slice2追加给slice2
	slice2 = append(slice2,slice2...)
	fmt.Println("slice2=",slice2) //slice2= [宋江 李逵 及时雨 宋江 李逵 及时雨]
	
	//将改变之后的值交给了slice3，故slice2没有变化
	slice3 := append(slice2,"hahaha","heiheihei")
	fmt.Println("slice2=",slice2) //slice2= [宋江 李逵 及时雨 宋江 李逵 及时雨]
	fmt.Println("slice3=",slice3) //slice3= [宋江 李逵 及时雨 宋江 李逵 及时雨 hahaha heiheihei]
}
