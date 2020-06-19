package main

import "fmt"

func main() {
	var arrs = [...]int{1,2,3,4,5,6}
	//方式一：常规遍历
	for i:=0;i<len(arrs);i++ {
		fmt.Printf("%d\t",arrs[i])
	}
	fmt.Println()
	
	//方式二：for-range遍历
	var arrStrs = [...]string{"张三","李四","王二"}
	for index,val:=range arrStrs{
		fmt.Printf("index=%d，val=%s\n",index,val)
	}
}


