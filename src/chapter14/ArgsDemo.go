package main

import (
	"fmt"
	"os"
)

//Args保管了命令行参数，第一个是程序名
func main() {
	fmt.Println("命令行参数有", len(os.Args))
	for idx, val := range os.Args {
		fmt.Printf("args[%v]=%v\n", idx, val)
	}
}
