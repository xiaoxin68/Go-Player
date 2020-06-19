package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	fileName1 := "a.txt"
	fileName2 := "D:/development/jetbrains/goland/workspace/src/Go-Player/hanru_package/a.txt"

	//判断是否为绝对路径
	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))

	//返回绝对路径
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2))

	//获取父目录
	fmt.Println(path.Join(fileName2, ".."))

	//创建目录:Mkdir只能创建一层文件夹，并且如果文件夹创建过了再次创建会报错
	/*err := os.Mkdir("D:/development/jetbrains/goland/workspace/src/Go-Player/hanru_package/bbb", os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("文件夹创建成功")*/

	//多层文件夹创建
	/*err := os.MkdirAll("D:/development/jetbrains/goland/workspace/src/Go-Player/hanru_package/cc/dd/ee", os.ModePerm)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("多层文件夹创建成功")*/

	//创建文件:如果文件不存在就创建，文件存在就不创建
	/*file,err := os.Create("D:/development/jetbrains/goland/workspace/src/Go-Player/hanru_package/bbb/a.txt")
	if err != nil {
		fmt.Println("err:",err)
		return
	}
	fmt.Println("文件创建成功:",file)*/

	//打开文件
	/*file, err := os.Open("./bbb/a.txt") //只读
	if err != nil {
		fmt.Println("文件打开失败:",err)
		return
	}
	fmt.Println("文件打开成功：",file)*/

	//OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。
	/*file, err := os.OpenFile("./bbb/a.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("文件打开失败:",err)
		return
	}
	fmt.Println("文件打开成功：",file)*/

	//关闭文件
	//file.Close()

	//删除文件或目录：Remove删除文件可直接删除，删除目录需要目录为空才可删除
	/*err := os.Remove("./bbb/a.txt")
	if err != nil {
		fmt.Println("删除文件失败:", err)
		return
	}
	fmt.Println("删除文件成功")*/

	//删除目录需要目录为空才可删除[remove ./cc: The directory is not empty.]
	//使用RemoveAll即使目录不为空也可直接删除
	err := os.RemoveAll("./cc")
	if err != nil {
		fmt.Println("删除目录失败:", err)
		return
	}
	fmt.Println("删除目录成功")
}
