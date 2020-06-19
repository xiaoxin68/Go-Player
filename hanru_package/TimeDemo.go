package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	t1 := time.Now()
	fmt.Println(t1)

	//获取指定时间
	t2 := time.Date(2019, 7, 10, 10, 20, 23, 0, time.UTC)
	fmt.Println(t2)

	//time->string
	s1 := t1.Format("2006-01-02 15:04:05")
	fmt.Println(s1)

	s2 := t2.Format("2006/1/2")
	fmt.Println(s2)

	//string->time
	s3 := "2019年07月10日 10时20分23秒"
	t3, err := time.Parse("2006年01月02日 15时04分05秒", s3)
	if err != nil {
		fmt.Println("转化失败")
	} else {
		fmt.Println(t3)
	}

	//时间戳
	fmt.Println(t1.Unix())
	fmt.Println(t1.UnixNano())

	//判断时间前后
	fmt.Println(t1.Before(t2))
	fmt.Println(t1.After(t2))
	fmt.Println(t2.Equal(t3))

	//输出具体的日期
	date, month, day := t1.Date()
	fmt.Println(date,month,day)
	fmt.Println(t1.Year())
	fmt.Println(t1.Month())
	fmt.Println(t1.Day())

	startDate, err := time.Parse("2006-01-02 15:04:05", "2020-05-03 19:32:15")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(startDate)

	d, err := time.Parse("2006-01-02", "2020-05-03")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
