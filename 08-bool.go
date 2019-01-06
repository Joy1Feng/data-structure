package main
import (
	"fmt"
	"unsafe"
)


func main(){
	var a bool = false
	fmt.Println("a=",a)
	fmt.Println("a的占用空间=",unsafe.Sizeof(a))
}