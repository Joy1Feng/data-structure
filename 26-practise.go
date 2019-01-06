package main
import (
	"fmt"
)


func main(){
	// var sum int = 0
	// for i := 1; i <= 100; i++ {
	// 	sum += i
	// 	if sum > 20 {
	// 		fmt.Println(i)
	// 		break
	// 	}
	// }
	var name string
	var pwd string
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s", "请输入用户名:")
		fmt.Scanln(&name)
		fmt.Printf("%s", "请输入密码:")
		fmt.Scanln(&pwd)
		if name == "joy" && pwd == "123" {
			fmt.Println("登录成功!")
			break
		} else {
			if i == 3 {
				fmt.Println("明天再来!")
			} else {
				fmt.Printf("你还有%d机会\n", 3 - i)
			}
		}
	}
}