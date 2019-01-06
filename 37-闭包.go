package main
import (
	"fmt"
	"strings"
)


// func myOutter () func (int) int {
// 	var num int = 10
// 	return func (num1 int) int {
// 		num += num1
// 		return num
// 	}
// }

func myOutter () func (string) string {
	var ext string = ".jpg"
	return func (str string) string {
		var newExt string
		if str != ext {
			newExt = ext
		} else {
			newExt = str
		}
		return newExt
	}
}

func main () {
	// f := myOutter()
	// fmt.Println(f(1))
	// fmt.Println(f(2))
	// fmt.Println(f(3))
	myRes := strings.HasSuffix(".jpg",".jpg")
	fmt.Println(myRes)
	f := myOutter()
	res := f(".jpg")
	fmt.Println(res)
}