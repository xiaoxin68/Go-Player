package main

import "testing"

func TestStore(t *testing.T)  {
	//创建一个monster实例
	monster := &Monster{
		Name:  "张三",
		Age:   12,
		Skill: "看书、游泳、学习",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误，希望=%v,实际=%v\n",true,res)
	}
	t.Logf("monster.Store()测试成功")
}

func TestReStore(t *testing.T)  {
	var monster =  &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore()错误，希望=%v,实际=%v\n",true,res)
	}
	if monster.Name != "张三" {
		t.Fatalf("monster.ReStore()错误，希望=%v,实际=%v\n","张三",res)
	}
	t.Logf("monster.ReStore()测试成功")
}
