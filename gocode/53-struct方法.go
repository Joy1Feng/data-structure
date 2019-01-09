package main
import (
	"fmt"
	"strconv"
)


type Student struct{
	Name string
}

func(a Student) myTest() {
	fmt.Println(a.Name)
}

type Person struct{
	Name string
	Age int
}

func (person Person) cal(num1 int, num2 int) string {
	sum := num1 + num2
	fmt.Println(sum)
	res := person.Name + strconv.Itoa(person.Age) + strconv.Itoa(sum)
	return res
}

type Box struct{
}

func (box Box) boxCal(num1 float64, num2 float64 , num3 float64) (sum float64){
	sum = num1 * num2 * num3
	return
}


func main() {
	student := Student{"joy"}
	student.myTest()
	person := Person{"tom", 23}
	res := person.cal(1, 2)
	fmt.Println(res)
	var num1 float64
	fmt.Printf("%s", "请输入num1")
	fmt.Scanln(&num1)
	var num2 float64
	fmt.Printf("%s", "请输入num2")
	fmt.Scanln(&num2)
	var num3 float64
	fmt.Printf("%s", "请输入num3")
	fmt.Scanln(&num3)
	fmt.Println(num1, num2, num3)
	box := Box{}
	res1 := box.boxCal(num1, num2, num3)
	fmt.Println(res1)
	re
}