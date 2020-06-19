package main

import (
	"fmt"
	"strconv"
)

//基本类型转string类型
func main() {
	var n1 int16 = 123
	var f1 float32 = 1.234
	var b1 bool = true
	var by1 byte = 'a'
	
	var str string
	
	//使用 fmt.Sprintf方式转换
	str = fmt.Sprintf("%d", n1)
	fmt.Printf("str type:%T,str = %v\n", str, str) //str type:string,str = 123
	
	str = fmt.Sprintf("%f", f1)
	fmt.Printf("str type:%T,str = %v\n", str, str) //str type:string,str = 1.234000
	
	str = fmt.Sprintf("%t", b1)
	fmt.Printf("str type:%T,str = %v\n", str, str) //str type:string,str = true
	
	str = fmt.Sprintf("%c", by1)
	fmt.Printf("str type:%T,str = %v\n", str, str) //str type:string,str = a
	
	//第二种方式：使用strconv里面的函数
	var str2 string
	//Itoa是FormatInt(i, 10) 的简写
	str2 = strconv.FormatInt(int64(n1), 10)            //返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
	fmt.Printf("str2 type:%T,str2 = %v\n", str2, str2) //str2 type:string,str2 = 123
	
	str2 = strconv.FormatFloat(float64(f1), 'f', 10, 64) //10“表示这个小数保留10位，64表示这个小数是float64
	fmt.Printf("str2 type:%T,str2 = %v\n", str2, str2)   //str2 type:string,str2 = 1.2339999676
	
	str2 = strconv.FormatBool(b1)
	fmt.Printf("str2 type:%T,str2 = %v\n", str2, str2) //str2 type:string,str2 = true
	
}
