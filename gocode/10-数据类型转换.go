package main
import (
	"fmt"
)


func main(){
	// var n1 int32 = 64
	// var n2 float64 = float64(n1)
	// fmt.Printf("%T\n", n1)
	// fmt.Printf("%T\n", n2)
	// fmt.Println(n2)
	var n1 float32 = 64.2
	var n2 int32 = int32(n1)
	fmt.Printf("%T\n", n1)
	fmt.Printf("%T\n", n2)
	fmt.Println(n2)
}