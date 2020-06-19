package main

import "fmt"

func main() {
	var mapLit map[string]int
	var mapAssign map[string]int
	mapLit = map[string]int{"one":1,"two":2}
	mapCreated := make(map[string]float32)
	mapAssign = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssign["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}