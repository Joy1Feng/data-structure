package main
import (
	"fmt"
)


func main() {
	var arr [2][3]int = [2][3]int{{1 ,2 ,3}, {4 ,5, 6}}
	fmt.Println(arr)
	for _, value := range arr {
		for _, value1 := range value {
			fmt.Println(value1)
		}
	}
}