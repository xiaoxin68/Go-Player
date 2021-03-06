package main

import "fmt"

func main() {
	//第一种方式
	for i := 1; i <= 10; i++ {
		fmt.Println("hello zxx")
	}
	
	//第二种方式
	j := 1
	for j <= 10 {
		fmt.Println("hello 尚硅谷")
		j++
	}
	
	//第三种方式
	k := 1
	for {
		if k <= 10 {
			fmt.Println("hello 小新新")
		}else {
			break
		}
		k++
	}
	
	//遍历字符串--传统方法【这种方法遍历字符串含中文字符会出现乱码】
	//h       e       l       l       o               å                       ä       º       ¬
	var s string = "hello 北京"
	for m := 0; m < len(s); m++ {
		fmt.Printf("%c \t",s[m])
	}
	fmt.Println()
	
	//解决方法，将字符转化成切片
	//h       e       l       l       o               北      京
	s2 := []rune(s)
	for m := 0; m < len(s2); m++ {
		fmt.Printf("%c \t",s2[m])
	}
	fmt.Println()
	
	//遍历字符串for-range方法【这种方法遍历字符串含中文字符不会出现乱码】
	var str string = "张晓信息想逆袭"
	for index ,val := range str{
		fmt.Printf("index=%d,val=%c\n",index,val)
	}
}
