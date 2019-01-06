package main
import (
	"fmt"
)


func main(){
	// var num1 int32 = 100
	// var f1 float32 = 21.32
	// var gender bool = false
	// var name string = "小王" 
	var str byte = 'a'
	var str1 string
	// str1 = fmt.Sprintf("%d", num1)
	// fmt.Printf("%T\n",str1)
	// fmt.Println(str1)
	
	// str1 = fmt.Sprintf("%t", gender)
	// fmt.Printf("%T\n",str1)
	// fmt.Println(str1)

	str1 = fmt.Sprintf("%c", str)
	fmt.Printf("%T\n",str1)
	fmt.Println(str1)
	

}