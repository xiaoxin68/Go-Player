package main

import (
	"fmt"
	"path/filepath"
)

//错误类型表示:只处理我们想要处理的错误类型
func main() {
	files, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		return
	}
	fmt.Println("matched files", files)
}
