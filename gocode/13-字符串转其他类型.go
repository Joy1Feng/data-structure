package main
import (
	"fmt"
	"strconv"
)


func main(){
	var str string = "false"
	var b bool
	b, _ = strconv.ParseBool(str)
	//函数有两个返回值, 只想获取第一个值, 忽略第二个值用 _ 可以忽略
	fmt.Printf("%T %v\n", b, b)
	var str1 string = "12345"
	var num int64
	num, _ = strconv.ParseInt(str1, 10, 64)
	fmt.Printf("%T %v\n", num, num)
	
}