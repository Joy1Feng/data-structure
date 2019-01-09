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
	cat.age = 23
	cat.gender = "m"
	fmt.Println(cat.name, cat.age, cat.gender)
	var cat1 Cat = cat
	cat1.name = "tom"
	fmt.Println(cat, cat1)
	// var cat1 *Cat = &cat
	// (*cat1).name = "tom"
	// fmt.Println(*cat1, cat)
	// fmt.Println(&(*cat1).name, &(*cat1).age, &(*cat1).gender)
}