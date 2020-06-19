package main

import (
	"fmt"
	"time"
)

func main() {
	//1、获取当前时间
	now := time.Now()
	//now的类型time.Time，now=2020-04-12 19:34:55.4525522 +0800 CST m=+0.038273801
	fmt.Printf("now的类型%T，now=%v\n",now,now)
	
	//2、返回时间点t对应的年、月、日。
	year, month, day := now.Date()
	fmt.Printf("%v年%v月%v日\n",year,month,day)
	fmt.Printf("%v年%v月%v日  %v:%v:%v\n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())

	//3、格式化日期时间：【注意】这个时间只能这么写
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:04:05"))
	
	//4、结合sleep使用时间常量
	i := 0
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Millisecond*1000)
		if i == 10 {
			break
		}
	}
	
	//5、Unix将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位秒）
	//UnixNano将t表示为Unix时间，即从时间点January 1, 1970 UTC到时间点t所经过的时间（单位纳秒）
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
}
