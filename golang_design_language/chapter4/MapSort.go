package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"sdewa":34,
		"alice":23,
		"sda":34,
		"sqwda":34,
	}
	var names []string
	for k,_ := range ages{
		names = append(names, k)
	}
	sort.Strings(names)
	for _,name := range names{
		fmt.Printf("%s\t%d\n",name,ages[name])
	}
}
