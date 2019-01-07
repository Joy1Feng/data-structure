package main
import (
	"fmt"
	"errors"
)


func myTest () {
	defer func() {
		// err := recover()
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误")
	}
}

func main () {
	myTest()
	fmt.Println("以后的代码")	
	err := readConf("config.ini1")
	if err != nil {
		panic(err)
	}
	fmt.Println("over")
}