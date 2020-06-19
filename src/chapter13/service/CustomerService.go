package service

import "../models"

type CustomerService struct {
	customers []models.Customer
	customerNum int
}
//返回*CustomerService的方法
func NewCustomerService() *CustomerService  {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	var customer = models.NewCustomer(1,"张三","男",20,"112","119@qq.com")
	customerService.customers = append(customerService.customers,*customer)
	return  customerService
}
//返回CustomerService.customers的方法
func (this *CustomerService) List()[]models.Customer {
	return this.customers
}
//1、添加客户
func (this *CustomerService)Add(customer models.Customer) bool {
	this.customerNum++
	customer.SetId(this.customerNum)
	this.customers = append(this.customers,customer)
	return  true
}
//2、删除客户
func (this *CustomerService)Delete(id int) bool {
	index := this.FindById(id)
	if  index == -1{
		return false
	}
	//如何从切片中删除一个元素
	this.customers = append(this.customers[:index],this.customers[index+1:]...)
	return true
}
//根据id找到对应的记录在切片中的位置
func (this *CustomerService)FindById(id int)  int{
	for i:=0;i< len(this.customers);i++  {
		if this.customers[i].GetId() == id{
			return  i
		}
	}
	return  -1
}



