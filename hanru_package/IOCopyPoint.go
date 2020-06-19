package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//断点续传
func main() {
	srcFile := "D:/development/jetbrains/goland/workspace/src/Go-Player/hanru_package/a.jpg"
	distFile := "copy_"+srcFile[strings.LastIndex(srcFile, "/")+1:]
	tempFile := distFile + "temp.txt"

	file1, err := os.Open(srcFile)
	handelError(err)
	file2, err := os.OpenFile(distFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	handelError(err)
	file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm) //可读可写
	handelError(err)

	defer file1.Close()
	defer file2.Close()
	defer file3.Close()

	//先读取临时文件中的数据，再seek
	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n1, err := file3.Read(bs)
	//handelError(err)
	countStr := string(bs[:n1])
	count, err := strconv.ParseInt(countStr, 10, 64)
	//handelError(err)
	//fmt.Println(count)

	//设置读，写的位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	data := make([]byte, 1024, 1024)
	n2 := -1            //读取的数据量
	n3 := -1            //写出的数据量
	total := int(count) //读取的总量

	//复制文件
	for {
		n2, err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件复制完毕:", total)
			file3.Close()
			//一旦复制完，就删除临时文件
			os.Remove(tempFile)
			break
		}
		n3, err = file2.Write(data[:n2])
		total += n3

		//将赋值的总量存储到临时文件中
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))

		fmt.Println("已经复制了", total, "字节数据")
		//模拟断电
		if total > 5000 {
			panic("断电啦")
		}
	}
}

func handelError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
