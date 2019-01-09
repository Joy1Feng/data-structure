package main
import (
	"fmt"
	"strconv"
)


type Account struct{
	Id int
	Pwd int
	Balance float64
}

func (account Account) login() bool {
	if account.Id == 1 && account.Pwd == 123456 {
		fmt.Println("登录成功")
		return true
	} else {
		fmt.Println("密码或帐号错误")
		return false
	}
}

func (account Account) withdraw(num float64) {
	if num > account.Balance {
		fmt.Println("取钱成功")
	} else {
		fmt.Println("取钱失败")
	}
}


func main() {
	var id string
	fmt.Printf("%s", "请输入id")
	fmt.Scanln(&id)
	var pwd string
	fmt.Printf("%s", "请输入pwd")
	fmt.Scanln(&pwd)
	newId, err1 := strconv.Atoi(id)
	newPwd, err2 := strconv.Atoi(pwd)
	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	account := Account{newId, newPwd, 0}
	res := account.login()
	if res {
		account.withdraw(12.2)
	} else {
		return
	}
}