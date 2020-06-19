package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//可以拷贝文件图片，音视频等
func main() {
	srcFile := "a.jpg"
	destFile := "./bbb/a_copy.jpg"
	n, err := copyFile1(srcFile, destFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("拷贝的总字节数为：", n)

	fmt.Println("================================")

	n2, err := copyFile2("a.jpg", "./bbb/a_copy2.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("拷贝的总字节数为：", n2)

	fmt.Println("================================")

	n3, err := copyFile3("a.jpg", "./bbb/a_copy3.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("拷贝的总字节数为：", n3)
}

//方式一：使用缓冲区的方式
func copyFile1(srcFile, destFile string) (int, error) {
	//打开源文件
	file, err := os.Open(srcFile)
	if err != nil {return 0, err}
	//打开目标文件
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {return 0, err}
	//关闭源文件,关闭目标文件
	defer file.Close()
	defer file2.Close()
	//拷贝数据
	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0 //总字节数
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {  //拷贝完毕
			break
		} else if err != nil {  //出错了
			return total, err
		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil
}

//方式二：使用io.Copy(dist,src)方法
//还有CopyN()复制src中的n个字节到dst
//还有CopyBuffer为指定一个buf缓存区，以这个大小完全复制
func copyFile2(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}

//方式三：使用ioutil包(一次性读写，不适用于文件过大的情况，文件过大可能造成内存溢出)
func copyFile3(srcFile, distFile string) (int, error) {
	file, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(distFile, file, os.ModePerm)
	if err != nil {
		fmt.Println("操作失败", distFile)
		return 0, err
	}
	return len(file), nil
}
