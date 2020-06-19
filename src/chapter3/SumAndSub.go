package main

import "fmt"

func main(){
	sum,sub:= sumAndSub(20, 30)
	fmt.Println("sum=",sum,"sub=",sub)
	
	result,_:=sumAndSub(23,20)
	fmt.Println("result=",result)
}

func sumAndSub(num1 int,num2 int)(int,int)  {
	sum := num1+num2
	sub := num1-num2
	return sum,sub
}
