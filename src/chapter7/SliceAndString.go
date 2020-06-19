package main

import "fmt"

func main() {
	//string的底层是byte数组，可以进行切片处理
	str := "hello@hubu.edu.cn"
	slice := str[6:]
	fmt.Printf("slice类型%T，slice=%v \n", slice, slice) //slice类型string，slice=hubu.edu.cn
	
	//注意：string本身不可变，那么如何修改string的元素呢
	//将string转化成[]byte或者[]rune,修改完成后再转化为string
	//byte按字节处理，一个汉字占三个字节；rune按字符处理
	bytes := []byte(str)
	bytes[0] = 'z' //注意：这里修改的字符不能是中文字符
	s1 := string(bytes)
	fmt.Println(s1)
	
	runes := []rune(str)
	//runes[0] = '北' //可以是中文字符
	s := string(runes)
	fmt.Println(s)
}
