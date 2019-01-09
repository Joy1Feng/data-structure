package main
import (
	"fmt"
	"os"
)

func PathExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	if !os.IsNotExist(err) {
		return true
	}
	return false
}



func main() {
	
	_, b := os.Stat("./hello.txt")
	fmt.Println(b)
	res := os.IsNotExist(b)
	fmt.Println(res)
}