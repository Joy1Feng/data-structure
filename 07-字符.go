package main
import (
	"fmt"
)


func main(){
	var c1 byte = 'a'
	var c2 byte = '0'
	//当我们直接输出byte值, 就是输出了对应的字符的码值
	fmt.Println("c1=", c1) //97
	fmt.Println("c2=", c2) //48
	fmt.Printf("c1=%c c2=%c",c1,c2)
}