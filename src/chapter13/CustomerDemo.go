package main

import (
	"./service"
	"./models"
	"fmt"
)

type customerView struct {
	key string
	loop bool
	customerService *service.CustomerService
}
//1、添加客户
func (this *customerView)add()  {
	fmt.Println("-----------------添加客户-----------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	customer := models.NewCustomer2(name,gender,age,phone,email)
	if this.customerService.Add(*customer) {
		fmt.Println("-----------------添加完成-----------------")
	}else {
		fmt.Println("-----------------添加失败-----------------")
	}
}
//2、修改客户
func (this *customerView)update()  {
	fmt.Println("-----------------修改客户信息-----------------")
}
//3、删除客户
func (this *customerView)delete()  {
	fmt.Println("-----------------删除客户-----------------")
	fmt.Println("请选择待删除的客户编号（-1退出）：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除（y/n）：")
	choise := ""
	fmt.Scanln(&choise)
	if choise == "y" ||choise=="Y" {
		if this.customerService.Delete(id){
			fmt.Println("-----------------删除完成-----------------")
		}else{
			fmt.Println("-----------------删除失败，输入的id不存在-----------------")
		}
	}
}
//4、客户列表
func (this *customerView)list()  {
	customers := this.customerService.List()
	fmt.Println("-----------------客户列表-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱\t")
	for i:=0;i< len(customers);i++  {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-----------------客户列表完成-----------------")
}
//5、客户列表
func (this *customerView)exit()  {
	fmt.Println("确认是否退出（y/n）：")
	for{
		fmt.Scanln(&this.key)
		if this.key == "y" || this.key=="Y"||this.key=="N"||this.key=="n" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出（y/n）：")
	}
	if this.key == "y" || this.key=="Y" {
		this.loop = false
	}
}
func (this *customerView)mainMenu()  {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1、添加客户")
		fmt.Println("                 2、修改客户")
		fmt.Println("                 3、删除客户")
		fmt.Println("                 4、客户列表")
		fmt.Println("                 5、退    出")
		fmt.Println("请选择（1-5）：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("请输入正确的选项。。。")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户信息管理系统")
}
func main() {
	cv := customerView{
		key:  "",
		loop: true,
	}
	cv.customerService = service.NewCustomerService()
	cv.mainMenu()
}
