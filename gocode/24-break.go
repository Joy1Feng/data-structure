package main
import (
	"fmt"
	"math/rand"
	"time"
)


func main(){
	//为了生成一个随机数, 还需要给rand设置一个种子
	rand.Seed(time.Now().Unix())
	//生成
	var count int = 0
	for {
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			fmt.Println(count)
			break
		}
	}
	// n := rand.Intn(100) + 1

	// fmt.Println(n)
}