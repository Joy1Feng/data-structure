package main
import (
	"fmt"
	"sort"
)


func main() {
	monsters := make([]map[string]string, 8)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "joy"
		monsters[0]["sex"] = "f"
	}
	newmonster := map[string]string{
		"name":"tom",
		"sex":"m",
	}
	monsters =  append(monsters, newmonster)
	fmt.Println(monsters, len(monsters))

	var student map[int]string = map[int]string{
		1 : "tom1",
		2 : "tom2",
		3 : "tom3",
		4 : "tom4",
	}
	fmt.Println(student)
	var keys []int
	for key, _ := range student {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, val := range keys {
		name, ok := student[val]
		if ok {
			fmt.Println(name)
		} else {
			fmt.Println("没找到")
		}
	}
	fmt.Println(keys)
}