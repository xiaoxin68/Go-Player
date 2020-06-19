package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts:=make(map[string]int)
	/*func NewScanner(r io.Reader) *Scanner
	NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines。*/
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}
	for line,n:=range counts{
		if n>1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}
