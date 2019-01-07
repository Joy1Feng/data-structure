package main
import (
	"fmt"
)

type Cat struct {
	name string
	age int
	gender string
}

func main() {
	var cat Cat
	cat.name = "joy"
	// cat.age = 23
	cat.gender = "m"
	fmt.Println(cat.name, cat.age, cat.gender)
}