package main
import (
	"fmt"
)

func main(){
// 	var str string = "hello"
// 	var ptr *string = &str
// 	fmt.Println(ptr,&str)
// 	fmt.Printf("%v",*ptr)
	var num int32 = 100
	fmt.Printf("%v\n", num)
	var ptr *int32 = &num
	*ptr = 200
	fmt.Printf("%v\n", num)
}