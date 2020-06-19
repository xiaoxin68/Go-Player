package main

import (
	"errors"
	"fmt"
)

//fmt:func Errorf(format string, a ...interface{}) error
//errors:func New(text string) error
func main() {
	//创建错误
	err1 := errors.New("自定义错误1")
	fmt.Printf("err1类型为%T,内容为：%v\n",err1,err1)

	err2 := fmt.Errorf("%v自定义的错误","张晓晓")
	fmt.Printf("err2类型为%T,内容为：%v\n",err2,err2)

	err := ageVertify(-10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("main end...")
}

func ageVertify(a int)error  {
	if a<=0 || a>=150 {
		return errors.New("年龄范围不合法")
	}
	return nil
}
