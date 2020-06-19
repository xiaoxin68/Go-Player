package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//统计英文、数字、空格和其他字符数量
type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	fileName := "C:/Users/zxx/Desktop/a.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	
	var count CharCount
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		runes := []rune(str) //为了兼容中文字符，可以将str转化成[]rune
		for _, val := range runes {
			//fmt.Printf("%v ",string(val))
			switch {
			case val >= 'a' && val <= 'z':
				fallthrough //穿透
			case val >= 'A' && val <= 'Z':
				count.ChCount++
			case val == ' ' || val == '\t':
				count.SpaceCount++
			case val >= '0' && val <= '9':
				count.NumCount++
			default:
				count.OtherCount++
				
			}
		}
	}
	fmt.Printf("字符个数=%v 数字个数=%v 空格个数=%v 其他字符个数=%v\n",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
