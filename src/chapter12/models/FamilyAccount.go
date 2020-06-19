package models

import "fmt"

type familyAccount struct {
	key     string  //保存用户输入的选项
	loop    bool    //是否循环
	balance float64 //账户余额
	money   float64 //每次收支金额
	note    string  //每次收支的说明
	flag    bool    //记录是否有收支的行为
	details string  //当有收支时，需要对details进行拼接处理
}

func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "金额\t账户金额\t收支金额\t说   明",
	}
}

//1、收支明细
func (fa *familyAccount) showDetails() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	if fa.flag {
		fmt.Println(fa.details)
	} else {
		fmt.Println("当前没有收支明细，来一笔吧！")
	}
}

//2、登记收入
func (fa *familyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&fa.money)
	fa.balance += fa.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&fa.note)
	//将这个情况拼接到detais变量
	fa.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", fa.balance, fa.money, fa.note)
	fa.flag = true
}

//3、登记支出
func (fa *familyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&fa.money)
	if fa.money > fa.balance {
		fmt.Println("余额的金额不足")
		return
	}
	fa.balance -= fa.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&fa.note)
	//将这个情况拼接到detais变量
	fa.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", fa.balance, fa.money, fa.note)
	fa.flag = true
}

//4、退出软件
func (fa *familyAccount) exit() {
	fmt.Println("你确定要退出吗？y/n")
	choise := ""
	for {
		fmt.Scanln(&choise)
		if choise == "y" || choise == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choise == "y" {
		fa.loop = false
	}
}
func (fa *familyAccount) MainMenu() {
	for {
		fmt.Println("-----------------家庭收支记账本-----------------")
		fmt.Println("                 1、收支明细")
		fmt.Println("                 2、登记收入")
		fmt.Println("                 3、登记支出")
		fmt.Println("                 4、退出软件")
		fmt.Println("请选择（1-4）：")
		fmt.Scanln(&fa.key)
		
		switch fa.key {
		case "1":
			fa.showDetails()
		case "2":
			fa.income()
		case "3":
			fa.pay()
		case "4":
			fa.exit()
		default:
			fmt.Println("请输入正确的选项。。。")
		}
		if !fa.loop {
			break
		}
	}
}
