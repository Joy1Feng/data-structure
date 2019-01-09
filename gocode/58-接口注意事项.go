package main
import (
	"fmt"
)

type Ainterface interface{
	Say()
}

type Student struct{

}
func (student Student) Say() {
	fmt.Println("i am say")
}



func main() {
	var student Student = Student{}
	var a Ainterface = student
	a.Say()
}