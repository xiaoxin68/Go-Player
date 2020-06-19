package main

import "fmt"

func main() {
	area, err := rectArea(10, -10)
	if err != nil {//如果出现了错误，可进一步判断错误类型
		fmt.Println(err)
		if ins, ok := err.(*rectError); ok {  //利用类型断言，将err转化成rectError类型。注意err实际上是一个指针，故加*
			//显示更加详细的错误信息
			if ins.heightError() {
				fmt.Printf("高度小于零错误，高度为：%d\n", ins.height)
			}
			if ins.widthError() {
				fmt.Printf("宽度小于零错误，宽度为：%d\n", ins.width)
			}
		}
		return
	}
	fmt.Println("面积为：", area)
}

type rectError struct {
	msg    string
	width  int
	height int
}
//自定义错误，需要实现Error() string接口
func (r *rectError) Error() string {
	return r.msg
}
//错误类型的判断函数
func (r *rectError) widthError() bool {
	return r.width < 0
}
//错误类型的判断函数
func (r *rectError) heightError() bool {
	return r.height < 0
}

func rectArea(width, height int) (int, error) {
	msg := ""
	if width < 0 {
		msg = "宽度小于0"
	}
	if height < 0 {
		if msg == "" {
			msg = "高度小于0"
		} else {
			msg += "，高度也小于0"
		}
	}
	if msg != "" {
		//注意：对于error类型返回，需要返回地址
		return 0, &rectError{msg, width, height}
	}
	return width * height, nil
}
