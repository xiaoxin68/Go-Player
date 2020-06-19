package main

import "fmt"

func main()  {
	fmt.Println(peachCount(1))
}
func peachCount(days int)int  {
	var res int
	if days>10 || days<0{
		res = -1
	} else if days == 10 {
		res = 1
	}else {
		res = (peachCount(days+1)+1)*2
	}
	return res
}
