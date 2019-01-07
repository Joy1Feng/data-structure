package main
import (
	"fmt"
)


func main(){
	// var str string = "中国重庆"
	// str = "中国成都"
	var str = "中国"
	var str2 = "重庆"
	var str3 = str + str2
	fmt.Println("str=", str)
	fmt.Println(str3)
	var str1 = `package main
	import (
		"fmt"
	)`
	fmt.Println(str1)
}