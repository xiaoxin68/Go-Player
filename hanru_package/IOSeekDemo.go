package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*func (*File) Seek
func (f *File) Seek(offset int64, whence int) (ret int64, err error)
Seek设置下一次读/写的位置。offset为相对偏移量，而whence决定相对位置：
0为相对文件开头，1为相对当前位置，2为相对文件结尾。它返回新的偏移量（相对开头）和可能的错误。
*/

func main() {
	fileName:="a.txt"
	file,err:=os.OpenFile(fileName,os.O_RDWR,os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//读写

	//读文件
	bs:=[]byte{0}
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4,io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs))

	//写文件
	file.Seek(0,io.SeekEnd)
	file.Write([]byte("哈哈哈"))
	file.WriteString("啦啦啦")
}
