package main
import (
	"fmt"
)


var age int = myTest()
func myTest () int {
	fmt.Println("test is using")
	return 90
}

func init () {
	fmt.Println("init is using")
}

func main () {
	fmt.Println(age)
}