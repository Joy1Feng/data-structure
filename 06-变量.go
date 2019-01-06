package main
// import "fmt"
// import "unsafe"
import (
	"fmt"
	"unsafe"
)


func main(){
	// var i int
	// i = 10
	// var i = 11.2
	// var (
	// 	name = "joy"
	// 	n1 = 23
	// 	n2 = 24
	// )
	// name := "joy"
	// var i int = 100
	// fmt.Println(i)
	// i = 10
	// fmt.Println(i)
	// i = 20
	// fmt.Println(i)	
	// // i = 1.1
	// fmt.Println(i)
	var i int = 1
	var j int = 2
	var r int = i + j
	fmt.Println(r)
	var country string = "中国"
	var city string = "重庆"
	var address string = country + city
	fmt.Println(address)
	// var z int8 = 1000
	// fmt.Println(z)
	var p = 10.1
	fmt.Printf("p %T",p)
	fmt.Println("")
	var n1 int64 = 100
	fmt.Printf("%T %d", n1, unsafe.Sizeof(n1))
	fmt.Println("")
	var f1 float64 = .123
	fmt.Printf("%T", f1)
	fmt.Println()
	fmt.Println(f1)
}