package main
import (
	"fmt"
)


func main(){
	// for i := 1; i <= 10; i++ {
	// 	if i % 2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	for i := 1; i <= 10; i++ {
		if i == 2 {
			goto label1
		}
		fmt.Println("hello1")
		fmt.Println("hello2")
		label1:
		fmt.Println("china")
	}
	
}