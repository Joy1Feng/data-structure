package main
import (
	"fmt"
)


type Interface interface{
	English()
}

type Student struct{
	Name string
	Age int
}
func (student Student) Study() {
	fmt.Printf("%d岁的%s正在学习", student.Age, student.Name)
}

type MidStudent struct{
	Student
	Score int
}
func (this MidStudent) Play() {
	fmt.Printf("%d岁的%s正在玩,得了%d分\n", this.Age, this.Name, this.Score)
}
func (this MidStudent) English() {
	fmt.Printf("%d岁的%s正在学英语,得了%d分\n", this.Age, this.Name, this.Score)
}


type Writer struct{
	Student
}
func (this Writer) Write() {
	fmt.Printf("%d岁的%s正在写\n", this.Age, this.Name,)
}
func (this Writer) English() {
	fmt.Printf("%d岁的%s正在学英语\n", this.Age, this.Name)
}

func main() {
	mid := MidStudent{
		Student{
			Name:"joy",
			Age : 23,
		}, 98,
	}
	mid.English()
	writer := Writer{
		Student{
			Name : "TOM",
			Age : 23,
		},
	}
	writer.English()
}
