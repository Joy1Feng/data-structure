package main
import (
	"fmt"
)


func main () {
	num1 := new(int)
	*num1 = 100
	fmt.Printf("%v %T %v\n", num1, num1, *num1)
}