package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //需要使用到mysql的init函数
)

func main() {
	db, err := sql.Open("mysql", "root:123456(127.0.0.1:3306)/taotao?charset=etf8")
	if err != nil {
		fmt.Println("错误信息：", err)
		return
	}
	fmt.Println("连接成功：", db)
}
