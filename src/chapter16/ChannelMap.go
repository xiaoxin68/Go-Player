package main

import "fmt"

func main() {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	
	m1 := make(map[string]string, 20)
	m1["name1"] = "张三"
	m1["name2"] = "李四"
	
	m2 := make(map[string]string, 20)
	m2["city1"] = "北京"
	m2["city2"] = "上海"
	
	mapChan <- m1
	mapChan <- m2
	
	fmt.Printf("mapChan[1]=%v,mapChan[2]=%v\n", <-mapChan, <-mapChan)
	//输出：mapChan[1]=map[name1:张三 name2:李四],mapChan[2]=map[city1:北京 city2:上海]
}
