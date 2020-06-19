package main

import (
	"fmt"
	"sort"
)

func main() {
	var classes = map[string][]string{
		"al":    {"datas"},
		"ca":    {"li"},
		"comp":  {"datas", "for", "compo"},
		"datas": {"dis"},
		"datab": {"datas"},
		"dis":   {"intro"},
		"net":   {"op"},
		"op":    {"datas", "compo"},
		"pro":   {"datas", "compo"},
	}
	for i, class := range topoSort(classes) {
		fmt.Printf("%d:\t%s\n", i+1, class)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	{
		visitAll = func(items []string) {
			for _, item := range items {
				if !seen[item] {
					seen[item] = true
					visitAll(m[item])
					order = append(order, item)
				}
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
