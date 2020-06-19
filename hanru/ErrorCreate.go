package main

import (
	"fmt"
	"math"
)

//自定义错误：实现Error() string接口
func main() {
	area, err := circleArea(-10.5)
	if err!=nil {
		fmt.Println(err)
		//判断错误的类型
		if ins,ok:=err.(*areaError);ok {
			fmt.Printf("\t错误类型为areaError,半径为%.2f\n",ins.radius)
		}
		return
	}
	fmt.Println("面积为：",area)
}

type areaError struct {
	msg    string
	radius float64
}

//自定义错误：实现Error() string接口
func (e *areaError) Error() string {
	return fmt.Sprintf("error:半径为%.2f,%s", e.radius, e.msg)
}

func circleArea(radius float64) (float64, error) {
	if radius <= 0 {
		return 0, &areaError{msg: "半径是非法的", radius: radius}
	}
	return math.Pi * radius * radius, nil
}
