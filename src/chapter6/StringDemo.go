package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1、统计字符串的长度，按字节func main()  {
	str := "hello北京"
	fmt.Println("str len = ", len(str)) //str len =  11
	
	//2、遍历字符串，解决中文字符的问题 []rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\t", r[i])
	}
	fmt.Println() //h       e       l       l       o       北      京
	
	//3、字符串转整数ParseInt、Atoi
	str2 := "23"
	//num1, _ := strconv.ParseInt(str2, 10, 64)
	num2, err := strconv.Atoi(str2) //Atoi是ParseInt(s, 10, 0)的简写。
	if err != nil {
		fmt.Println("转化错误")
	} else {
		fmt.Printf("类型为%T，值为%v\n", num2, num2) //类型为int，值为23
	}
	
	//4、整数转字符串:Itoa
	num := 22
	s := strconv.Itoa(num)           //Itoa是FormatInt(i, 10) 的简写
	fmt.Printf("类型为%T，值为%v\n", s, s) //类型为string，值为22
	
	//5、字符串转[]byte:[]byte(str)
	bytes := []byte(str)
	fmt.Printf("类型为%T，值为%v\n", bytes, bytes) //类型为[]uint8，值为[104 101 108 108 111 229 140 151 228 186 172]
	
	//6、[]byte转字符串:string([]byte{104,101})
	ss := string([]byte{104, 101, 108, 108, 111, 229, 140, 151, 228, 186, 172})
	sss := string(bytes)
	fmt.Printf("ss=%v,sss=%v\n", ss, sss) //ss=hello北京,sss=hello北京
	
	//7、十进制转其他进制
	var num3 int64 = 123
	fmt.Printf("%d对应的二进制是：%v\n", num3, strconv.FormatInt(num3, 2))
	fmt.Printf("%d对应的八进制是：%v\n", num3, strconv.FormatInt(num3, 8))
	fmt.Printf("%d对应的十六进制是：%v\n", num3, strconv.FormatInt(num3, 16))
	
	//8、查找某字符串是否包含指定的子串
	b := strings.Contains("hello", "llo") //true
	fmt.Println(b)
	
	//9、查号一个字符串有几个指定的子串
	count := strings.Count("hellomenomasom", "om") //3
	fmt.Println(count)
	
	//10、区分大小写的比较和不区分大小写的比较
	fmt.Println("abc" == "AbC")                  //false
	fmt.Println(strings.EqualFold("abc", "AbC")) //true
	
	//11、子串sep在字符串s中第一次出现的位置，不存在则返回-1。
	fmt.Println(strings.Index("helllloomllomlloll", "llo")) //4
	
	//12、子串sep在字符串s中最后一次出现的位置，不存在则返回-1
	fmt.Println(strings.LastIndex("helllloomllomlloll", "llo")) //13
	
	//13、返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
	//func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("helllloomllomlloll", "llo", "小", 2))  //hell小om小mlloll
	fmt.Println(strings.Replace("helllloomllomlloll", "llo", "小", 0))  //helllloomllomlloll
	fmt.Println(strings.Replace("helllloomllomlloll", "llo", "小", -1)) //hell小om小m小ll
	
	//14、按照指定的分割方法将字符串分割为数组
	str3 := "hello|haha|hahma|lalla"
	split := strings.Split(str3, "|")
	for i := 0; i < len(split); i++ {
		fmt.Printf("%s\t", split[i]) //	hello   haha    hahma   lalla
	}
	fmt.Println()
	
	//15、将字符串的字母进行大小写转换
	str4 := "fhgAdasJFsSA"
	fmt.Println(strings.ToLower(str4)) //fhgadasjfssa
	fmt.Println(strings.ToUpper(str4)) //FHGADASJFSSA
	fmt.Println(str4)                  //fhgAdasJFsSA
	
	//16、将字符串两边的空格去掉
	str5 := "  214 gfsd 343    "
	fmt.Println(strings.TrimSpace(str5)) //214 gfsd 343
	fmt.Println(str5)                    //  214 gfsd 343
	
	//17、将字符串两边指定的字符去掉:返回将s前后端所有cutset包含的utf-8码值都去掉的字符串
	str6 := "!!! 214 gfsd!!! 343!!!!"
	fmt.Println(strings.Trim(str6, "!"))   // 214 gfsd!!! 343
	fmt.Println(strings.Trim(str6, "!13")) // 214 gfsd!!! 34
	fmt.Println(str6)                      //!!! 214 gfsd!!! 343!!!!
	
	//18、将字符串左边指定的字符去掉
	fmt.Println(strings.TrimLeft(str6, "!")) // 214 gfsd!!! 343!!!!
	
	//19、将字符串右边边指定的字符去掉
	fmt.Println(strings.TrimRight(str6, "!")) //!!! 214 gfsd!!! 343
	
	//20、判断字符串是否以什么开头
	fmt.Println(strings.HasPrefix("hello", "he")) //true
	
	//21、判断字符串是否以什么结尾
	fmt.Println(strings.HasSuffix("world", "dd")) //false
}
