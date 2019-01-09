package main
import (
	"fmt"
)

type Person struct{
	Name string
	Age int
}

func (person Person) say() {
	fmt.Println(person.Name,person.Age)
}

type Student struct{
	Person
	Score int
}
func (student Student) Study() {
	fmt.Println("i am study")
}


type Graduate struct{
	g Student
}

func main() {
	// person := Person{"joy", 23}
	// person.say()
	// var student Student
	// student.Person.Name = "joy1"
	// // student.Name = "joy"
	// student.Age = 23
	var student Student
	student.Name = "joy2"
	student.Age = 23
	student.Score = 90
	student.say()
	student.Study()
	fmt.Println(student.Score)
	var g Graduate
	g.g.Name = "tom"
	fmt.Println(g)
}