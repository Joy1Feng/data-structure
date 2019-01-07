package main
import (
	"fmt"
)


func main(){
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("hello, world")
	// }
	// j := 1
	// for j <= 10 {
	// 	fmt.Println("hello, world")
	// 	j ++
	// }
	// for {
	// 	if j <= 10 {
	// 		fmt.Println("hello, world")
	// 	} else {
	// 		break
	// 	}
	// 	j ++
	// }
	var str string = "china是一个大国家"
	str1 := []rune(str) 
	for index, val := range str1 {
		fmt.Printf("%c %d\n", val, index)
	}
}