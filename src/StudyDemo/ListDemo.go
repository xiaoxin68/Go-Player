package main

import (
	"container/list"
	"fmt"
)

func main() {
	/*1) 通过 container/list 包的 New() 函数初始化 list
	变量名 := list.New()
	
	2) 通过 var 关键字声明初始化 list
	var 变量名 list.List*/
	
	l := list.New()
	
	//双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。
	// 尾部添加
	l.PushBack("canon") //PushBack将一个值为v的新元素插入链表的最后一个位置，返回生成的新元素。
	// 头部添加
	l.PushFront(67) //PushFront将一个值为v的新元素插入链表的第一个位置，返回生成的新元素。
	
	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)
	// 使用
	l.Remove(element)
	
	//遍历列表——访问列表的每一个元素
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
