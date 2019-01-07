package main
import (
	"fmt"
	"strconv"
)


func main() {
	// // var dic map[int]string
	// dic := make(map[int]string, 2)
	// // dic = make(map[int]string, 2)
	// dic[0] = "joy"
	// dic[1] = "lulu"

	// heros := map[string]string{
	// 	"hrero1": "joy",
	// 	"hrero2": "joy",
	// 	"hrero3": "joy",
	// }
	// heros["hero4"] = "lulu"
	// fmt.Println(dic, heros)
	var student map[int]string
	student = make(map[int]string)
	for i := 1; i < 6; i++ {
		student[i] = "tom" + strconv.Itoa(i)
	}
	fmt.Println(student)
	for index, val := range student {
		fmt.Println(index, val)
		// val, ok := student[index]
		// if ok {
		// 	fmt.Println(val)
		// } else {
		// 	fmt.Println("不存在")
		// }
	}
}