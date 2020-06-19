package main

import "fmt"

type tree struct {
	value        int
	left, rigrht *tree
}

func sort(values []int) {
	var root *tree
	//将[]int构造成一个排序二叉树
	for _, v := range values {
		root = add2(root, v)
	}
	//中序遍历二叉树
	appendValues(values[:0], root)
}

//append元素按顺序追加到values
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.rigrht)
	}
	return values
}
//构造排序二叉树
func add2(t *tree, value int) *tree {
	if t == nil {
		return &tree{value: value}
	}
	if value < t.value {
		t.left = add2(t.left, value)
	} else {
		t.rigrht = add2(t.rigrht, value)
	}
	return t
}
func main() {
	var values = []int{3,5,8,6,2,4}
	sort(values)
	fmt.Println(values)
}
