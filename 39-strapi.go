package main
import (
	"fmt"
	"strconv"
)

func myStr (str string) {
	//循环遍历数组
	strLen := len([]rune(str))
	fmt.Println(strLen)
	for i := 0; i < strLen; i++ {
		fmt.Printf("%c\n", []rune(str)[i])
	}
}

func main () {
	myStr("中国china")
	res := len("hello中国")//字符串的额长度
	fmt.Println(res)
	val, err := strconv.Atoi("10000")//字符串转数字
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
	
}