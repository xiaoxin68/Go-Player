package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	
	tip1 := "genji is a ninja"
	fmt.Println(len(tip1))  //16
	
	tip2 := "忍者"
	fmt.Println(len(tip2)) //3
	
	//如果希望按习惯上的字符个数来计算，就需要使用 Go 语言中 UTF-8 包提供的 RuneCountInString() 函数，统计 Uncode 字符数量。
	fmt.Println(utf8.RuneCountInString("忍者")) //2
	fmt.Println(utf8.RuneCountInString("龙忍出鞘,fight!"))//11
	
}
