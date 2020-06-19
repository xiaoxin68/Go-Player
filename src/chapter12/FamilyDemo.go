package main

import (
	"./models"
)

func main() {
	var account = models.NewFamilyAccount()  //返回的是指针变量
	(*account).MainMenu()
	
	//改进：使用软件前有一个登录功能，只有正确输入用户名和密码才能访问
	
}
