package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//fmt包
	/*var x int
	var y float32
	fmt.Scanln(&x,&y)
	fmt.Println(x,y)

	fmt.Scanf("%d,%f",&x,&y)
	fmt.Println(x,y)*/

	//bufio包
	fmt.Println("请输入一个字符型：")
	reader := bufio.NewReader(os.Stdin)
	s,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("读到的字符串：",s)
}
