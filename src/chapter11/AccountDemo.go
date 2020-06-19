package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

func (account *Account)Deposite(money float64,pwd string)  {
	if pwd != account.Pwd{
		fmt.Println("密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("金额不正确")
		return
	}
	account.Balance += money
	fmt.Println("存款成功")
}

func (account *Account)WithDraw(money float64,pwd string)  {
	if pwd != account.Pwd{
		fmt.Println("密码不正确")
		return
	}
	if money <= 0 || money > account.Balance{
		fmt.Println("金额不正确")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功")
}

func (account *Account)Query(pwd string)  {
	if pwd != account.Pwd{
		fmt.Println("密码不正确")
		return
	}
	fmt.Printf("你的账号为：%x，账户余额为：%v\n",account.AccountNo,account.Balance)
}
func main() {
	var account = Account{
		AccountNo: "10011",
		Pwd:       "1234",
		Balance:   345.89,
	}
	account.Query("1234")
	account.Deposite(25.89,"1234")
	account.WithDraw(156.09,"1234")
	account.WithDraw(300.09,"1234")
}
