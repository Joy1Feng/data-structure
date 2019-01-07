package main
import (
	"fmt"
	"math/rand"
	"time"
)


func cal(ptr *[3]int32) {
	(*ptr)[0] = 10

	fmt.Println(*ptr)
}


func bornint() int{
	rand.Seed(time.Now().UnixNano())
	res10 := rand.Intn(100)
	return res10
}


func main() {
	//var intarr [5]int
	// intarr[0] = 10
	// intarr[1] = 20
	// intarr[2] = 30
	// intarr[3] = 40
	// intarr[4] = 50

	// var intarr [3]int = [3]int{1, 2, 3}

	//var intarr = [3]int{1, 2, 3}

	//var intarr = [...]int{1, 2, 3, 4}
	//fmt.Printf("%T", intarr)

	var intarr = [...]int{0:100, 2:200, 1:300}
	fmt.Println(intarr)
	// for i := 0; i < len(intarr); i++ {
	// 	fmt.Println(intarr[i])
	// }
	for index, value := range intarr {
		fmt.Println(index, value)
	}
	var arrfloat [3]float64
	fmt.Println(arrfloat)

	var arrint32 [3]int32
	// var ptrarrin32 *[3]int32 = &arrint32
	cal(&arrint32)
	fmt.Println(arrint32)

	var mystr [26]byte
	for i := 0; i < len(mystr); i++ {
		mystr[i] = 'A' + byte(i)
	}
	for i := 0; i < len(mystr); i++ {
		fmt.Printf("%c\n", mystr[i])
	}
	maxarrint  := [...]int{1, 2, 3, 4, 0}
	maxint := maxarrint[0]
	for i := 0; i < len(maxarrint); i++ {
		if maxarrint[i] > maxint {
			maxint = maxarrint[i]
		}
	}
	sum := maxarrint[0]
	for i := 1; i < len(maxarrint); i++ {
		sum += maxarrint[i]
	}
	var sum1 int 
	for _, value := range maxarrint {
		sum1 += value
	}
	fmt.Println(maxint, sum, sum1)

	var sortarrint [5]int
	for index, _ := range sortarrint {
		sortarrint[index] = bornint()
	}
	fmt.Println(sortarrint)
	for index, _ := range sortarrint {
		fmt.Println(sortarrint[len(sortarrint)-index-1])
	}
	
}