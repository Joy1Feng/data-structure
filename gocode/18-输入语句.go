package main
import (
	"fmt"
)


func main(){
	var name string

	var age int32
	// fmt.Printf("%s", "请输入姓名:")
	// fmt.Scanln(&name)
	// fmt.Printf("%s", "请输入年龄:")
	// fmt.Scanln(&age)
	// fmt.Printf("名字是%v年龄是%d\n\n", name, age)
	fmt.Println("姓名, 年龄, 请使用,隔开")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("名字是%v年龄是%d\n\n", name, age)
}
