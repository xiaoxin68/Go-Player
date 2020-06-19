package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str1 string = "hello 张晓新" //由于编码是utf-8，所以不会出现中文乱码问题
	fmt.Println("str1=", str1)
	fmt.Printf("str1占用字节数：%d \n", unsafe.Sizeof(str1)) //str1占用字节数：16
	
	//string内容是不可变的
	
	//可以识别转义字符
	str2 := "abc\tdef" //str2= abc       def
	fmt.Println("str2=", str2)
	
	//反引号，使字符串以原生形式输出,包括代码换行转义，可以防止攻击
	str3 := `package main

			import "fmt"
			
			func main()  {
				fmt.Println("张三\t李四")
				fmt.Println("张三\n李四")
				fmt.Println("张三\\李四")
				fmt.Println("张三\"李四")
				fmt.Println("张三\r李四")  //表示从当行的最前面输出，替换掉前面的内容
			}
			`
	fmt.Println("str3=", str3)
	
	//字符串的拼接
	str4 := "hello" + "world"
	str4 += "\t张晓新"
	str4 += `\t小新新`
	fmt.Println("str4=", str4)
	
	//如果字符串过长:注意+要保留在上一行，否则会出问题
	str5 := "hello" + "world"+"hello" + "world"+"hello" + "world"+
	"hello" + "world"+"hello" + "world"+"hello" + "world"+
	"hello" + "world"+"hello" + "world"
	fmt.Println("str5=", str5)
}
