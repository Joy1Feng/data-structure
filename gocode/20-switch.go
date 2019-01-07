package main
import (
	"fmt"
)


func main(){
	// var a string = "l"
	var num int = 7
	// switch num + 1 {
	// 	case 1, 2:
	// 		fmt.Println(1)
	// 	// case 2:
	// 	// 	fmt.Println(2)
	// 	case 3:
	// 		fmt.Println(3)
	// 	default:
	// 		fmt.Println(4)
	// }
	switch {
		case num == 0:
			fmt.Println(0)
		case num == 1:
			fmt.Println(1)
		default:
			fmt.Println(3)
	}	

}