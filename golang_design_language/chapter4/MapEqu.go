package main

import "fmt"

func main() {
	fmt.Println(mapEqu(map[string]int{"A": 0}, map[string]int{"B": 42}))
}

func mapEqu(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		//if xv!=y[k] {//注意一定不要这么写，否则fmt.Println(mapEqu(map[string]int{"A":0}, map[string]int{"B":42}))执行结果为true
		if yv, ok := y[k]; !ok || yv != xv {

			return false
		}
	}
	return true
}
