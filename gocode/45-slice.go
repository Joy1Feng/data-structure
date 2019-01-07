package main
import (
	"fmt"
)


func main() {
	//第一种创建方式
	var intarr = [...]int{1, 22, 33, 44, 55, 66, 77}
	slice := intarr[2:5]
	fmt.Printf("%T\n", slice)
	fmt.Println(slice, cap(slice), &slice[0])
	
	//第二种创建方式
	var int32arr = make([]int, 5, 10)
	fmt.Println(int32arr)

	//第三种创建方式
	var mystr []string = []string{"lulu", "tom", "joy"}
	fmt.Println(mystr, &mystr[0])
	mystr1 := append(mystr, "中国", "重庆")
	fmt.Println(mystr, &mystr1[0])
	// num10 := new(int)
	// *num10 = 10
	// fmt.Println(*num10)
}