package main
import (
	"fmt"
)


func main(){
	// label1:
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		label1:
		for z := 1; z < 10; z++ {
			if z == 2 {
				break label1
			}
			fmt.Println(z)
		}
	}
	fmt.Println("over")
}