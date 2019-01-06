package main
import (
	"fmt"
)


func cal (num int32) {
	if num > 2 {
		num--
		cal(num)
	} else {
		fmt.Println(num)
	}
	
}

func main(){
	cal(4)
}