package util

import "fmt"

type FamilyAccount struct {
	key  string //  用户输入选项
	loop bool

	//定义账户的余额
	balance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	//
	flag bool
	// 收支的详细使用字符串来记录，当有收支时，只需要对 Details 进行拼接处理即可
	details string
}

// 写一个工厂模式的构造方法，返回一个 *FamilyAccount 实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说  明",
	}
}

// 结构体绑定方法
func (this *FamilyAccount) MainMenu() {
	// 循环显示
	for {
		fmt.Println("---------家庭收支记账软件--------------")
		fmt.Println("          1 收入明细         ")
		fmt.Println("          2 登记收入         ")
		fmt.Println("          3 登记支出         ")
		fmt.Println("          4 退出收软件         ")
		fmt.Println("请选择(1-4)：")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.add()
		case "3":
			this.sub()
		case "4":
			this.exist()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop {
			break
		}
	}
}

func (this *FamilyAccount) exist() {
	fmt.Println("确定要退出吗 y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("请输入正确的选项")
	}
	if choice == "y" {
		this.loop = false
	}
}

func (this *FamilyAccount) sub() {
	fmt.Println("登记支出")
	fmt.Scanln(&this.money)
	if this.balance < this.money {
		fmt.Println("余额不足")
		//return
	}
	this.balance -= this.money // 修改收入
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出 \t%v \t%v \t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) add() {
	fmt.Println("收入金额 ")
	fmt.Scanln(&this.money)
	this.balance += this.money // 修改收入
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入 \t%v \t%v \t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) showDetails() {
	fmt.Println("显示收入明细")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支，请录入")
	}
}
