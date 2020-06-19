package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)  //unicode字符数量
	var uftlen [utf8.UTFMax+1]int //utf-8编码的长度
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for{
		r,n,err := in.ReadRune()
		if err==io.EOF {
			break
		}
		if err!=nil {
			fmt.Fprintf(os.Stderr,"charcount:%v\n",err)
		}
		if r == unicode.ReplacementChar && n==1 {
			invalid++
			continue
		}
		counts[r]++
		uftlen[n]++
	}
	fmt.Printf("rune\tcount\t")
	for c,n := range counts{
		fmt.Printf("%q\t%d\n",c,n)
	}
	for i,n := range uftlen{
		if i>0 {
			fmt.Printf("%q\t%d\n",i,n)
		}
	}
	if invalid > 0{
		fmt.Printf("\n%d invalid UTF-8 characters\n",invalid)
	}
}
