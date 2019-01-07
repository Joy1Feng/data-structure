package main
import (
	"fmt"
	"strconv"
)

func main(){
	var num int32 = 100
	var str string
	str = strconv.FormatInt(int64(num), 10)
	fmt.Println(str)
	var f1 float64 = 12.2
	str = strconv.FormatFloat(f1, 'f', 10, 64)
	fmt.Println(str)
}
