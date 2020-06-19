package main

import (
	"testing"
)

//运行：go test -v  无论正确还是错误都输出日志
//运行：go test  正确不输出日志，错误输出日志
//测试单个文件一定要带上被测试的原文件:go test -v Cal_test.go Cal.go
//测试单个方式：go test -v -test.run TestAddUpper

//注意：【文件名】go test 命令，会自动读取源码目录下面名为 *_test.go 的文件，生成并运行测试用的可执行文件
//注意：【函数名称】单元测试源码文件可以由多个测试用例组成，每个测试用例函数需要以Test为前缀，例如： func TestXXX( t *testing.T )
//TestXXX后面的XXX第一个字母一定要大写
func TestAddUpper(t *testing.T)  {
	res := AddUpper(10)
	if res !=55{
		//Fatalf:相当于在调用 Logf 之后调用 FailNow
		//FailNow:将当前的测试函数标识为“失败”，并停止执行该函数
		t.Fatalf("addUpper(10) 执行错误，期望值=%v，实际值=%v\n",55,res)
	}
	//如果正确：记录日志
	//Log 使用与 Printf 相同的格式化语法对它的参数进行格式化， 然后将格式化后的文本记录到错误日志里面
	t.Logf("test addUpper 执行正确\n")
}
