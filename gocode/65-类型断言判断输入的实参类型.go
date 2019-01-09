package main
import (
	"fmt"
)


type Student struct{
	Name string
}


//编写一个函数,可以判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for index, val := range items {
		index++
		switch val.(type) {
			case bool :
				fmt.Printf("第%v个参数是bool类型, 值是%v\n", index, val)
			case float32 :
				fmt.Printf("第%v个参数是float32类型, 值是%v\n", index, val)
			case float64 :
				fmt.Printf("第%v个参数是float64类型, 值是%v\n", index, val)
			case int, int32, int64 :
				fmt.Printf("第%v个参数是int类型, 值是%v\n", index, val)
			case string :
				fmt.Printf("第%v个参数是string类型, 值是%v\n", index, val)
			case Student :
				fmt.Printf("第%v个参数是Student类型, 值是%v\n", index, val)
			case *Student :
				fmt.Printf("第%v个参数是*Student类型, 值是%v\n", index, val)
			default :
			fmt.Printf("第%v个参数是类型不确定, 值是%v\n", index, val)
		}
	}
	
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300
	student := Student{
		Name : "joy",
	}
	student1 := &Student{
		Name : "tom",
	}
	TypeJudge(n1, n2, n3, name, address, n4, student, student1)
}