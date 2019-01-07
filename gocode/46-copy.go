package main
import (
	"fmt"
)


func main() {
	var intarr []int = []int{1, 2, 3, 4}
	var newintarr []int = make([]int, 1)
	newintarr[0] = 10
	copy(newintarr, intarr)
	fmt.Println(intarr, newintarr, &intarr[0], &newintarr[0])

	var str string = "hello中国"
	// arrbyte := []byte(str)
	arrbyte := []rune(str)
	arrbyte[0] =  'H'
	arrbyte[5] =  '重'
	str = string(arrbyte)
	fmt.Println(str)
}