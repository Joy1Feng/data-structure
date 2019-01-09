package main
import (
	"fmt"
	"encoding/json"
)


type Student struct {
	Name string
	Age int
}

type Stu Student


type integer int

type Monster struct {
	Name string `json:"name"` //`json:"name"`就是struct tag
	Age int `json:"age"`
	Skill string `json:"skill"`
}


func main() {
	var stu1 Student
	var stu2 Stu
	stu2 = Stu(stu1)
	fmt.Println(stu1, stu2)
	// var i integer = 10
	// var j int = 20
	// j = i
	m1 := Monster{"tom", 23, "play"}
	jsonStr, _ := json.Marshal(m1)
	fmt.Println(m1, string(jsonStr))
}