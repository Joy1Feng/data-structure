package main
import (
	"fmt"
	// "strconv"
	"strings"
)

// func myStr (str string) {
// 	//循环遍历数组
// 	strLen := len([]rune(str))
// 	fmt.Println(strLen)
// 	for i := 0; i < strLen; i++ {
// 		fmt.Printf("%c\n", []rune(str)[i])
// 	}
// }

func main () {
	// myStr("中国china")
	// res := len("hello中国")//字符串的长度
	// fmt.Println(res)
	// val, err := strconv.Atoi("10000")//字符串转数字
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(val)
	// }
	
	//数字转字符串
	// var num int = 32
	// str := strconv.Itoa(num)
	// fmt.Printf("%v %T\n", str, str)

	//字符串转byte, byte转字符串
	// var str string = "hello中国"
	// var bytes = []byte(str)
	// fmt.Printf("%v %T\n", bytes, bytes)
	// str = string(bytes)
	// fmt.Printf("%v %T\n", str, str)

	//判断子串是否在字符串中
	// b := strings.Contains("hello", "he")
	// fmt.Println(b) 
	
	//统计指定子串出现的次数
	// res := strings.Count("helloll", "ll")
	// fmt.Println(res)

	//不区分大小写
	// res := strings.EqualFold("abc", "ABc")
	// fmt.Println(res)

	//返回第一次出现的索引
	// res := strings.LastIndex("hello world", "o")
	// fmt.Println(res)

	//替换子串
	// res := strings.Replace("hello world", "l", "L", 2)
	// fmt.Println(res)

	//按照指定的标记拆分为字符串数组
	//strArr := strings.Split("hello,world,china", ",")
	//fmt.Println(strArr[0])
	
	//去掉字符串左右两边指定字符
	// res := strings.Trim("! hello!", "!")
	// fmt.Println(res)

	//是否以指定的字符串开头
	res := strings.HasPrefix("https://www.baidu.com", "https")
	fmt.Println(res)
}