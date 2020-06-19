package main

import (
	"flag"
	"fmt"
	"strings"
)

//标识的名字 默认值 非法参数或help时的提示信息
var n = flag.Bool("n",false,"omit trialing newLine")  //n是指针
var sep = flag.String("s", " ","separator")
func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*sep))
	if !*n {
		fmt.Println()
	}
}
/*go run PtrDemo.go a bd def
a bd def*/

/*go run PtrDemo.go -s / a bc def
a/bc/def*/

/*go run PtrDemo.go -n a bc def
a bc def$*/

/*go run PtrDemo.go -help
Usage of C:\Users\zxx\AppData\Local\Temp\go-build239792062\b001\exe\PtrDemo.exe:
-n    omit trialing newLine
-s string
separator (default " ")*/


