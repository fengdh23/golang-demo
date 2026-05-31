package main

import "fmt"

func main() {
	key := "" //  用户输入选项
	loop := true

	//定义账户的余额
	balance := 100.0
	// 每次收支的金额
	money := 0.0
	// 每次收支的说明
	note := ""
	//
	flag := false
	// 收支的详细使用字符串来记录，当有收支时，只需要对 Details 进行拼接处理即可
	var details = "收支\t账户金额\t收支金额\t说  明"

	for {

		fmt.Println("---------家庭收支记账软件--------------")
		fmt.Println("          1 收入明细         ")
		fmt.Println("          2 登记收入         ")
		fmt.Println("          3 登记支出         ")
		fmt.Println("          4 退出收软件         ")
		fmt.Println("请选择(1-4)：")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("显示收入明细")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支，请录入")
			}
		case "2":
			fmt.Println("收入金额 ")
			fmt.Scanln(&money)
			balance += money // 修改收入
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入 \t%v \t%v \t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("登记支出")
			fmt.Scanln(&money)
			if balance < money {
				fmt.Println("余额不足")
				break
			}
			balance -= money // 修改收入
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出 \t%v \t%v \t%v", balance, money, note)
			flag = true
		case "4":
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
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项")
		}
		if !loop {
			break
		}
	}
}
