package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

//遍历文件夹
func main() {
	dirName := "D:/development/jetbrains/goland/workspace/src/Go-Player"
	listFiles(dirName, 0)
}

func listFiles(dirName string, level int) {
	s := "|-"
	for i := 0; i < level; i++ {
		s = "| " + s
	}
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirName + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() {
			//递归调用此方法
			listFiles(filename, level+1)
		}
	}
}
