package main
import (
	"fmt"
)


func main(){
	// var i int = 1
	// var count int
	// for i <= 100 {
	// 	if i % 9 == 0 {
	// 		fmt.Println(i)
	// 		count++
	// 	}
	// 	i++
	// }
	// fmt.Println(count)
	var i int = 6
	var j int = 0
	for j <= i {
		fmt.Printf("%d + %d = %d\n", j, i - j, i)
		j++
	}
}