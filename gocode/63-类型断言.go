package main
import (
	"fmt"
)


type Point struct{
	Num1 int
	Num2 int
}

func main() {
	//创建一个空的interface
	var Interface interface{}
	fmt.Println(Interface)
	//创建一个a点
	var a Point = Point{1, 2}
	//将a点赋值给空接口
	Interface = a
	var b Point
	//创建一个b点,并进行类型断言、
	b = Interface.(Point)
	//把a赋值给b
	fmt.Println(Interface, b)

	var Interface1 interface{}
	var x float64 = 9.9
	Interface1 = x
	var y float64
	y, ok := Interface1.(float64)
	if ok {
		fmt.Println("装换成功")
	} else {
		fmt.Println("转换失败")
	}
	fmt.Println("============")
	fmt.Printf("%T %v\n", y, y)
}