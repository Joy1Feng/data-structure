package main
import (
	"fmt"
	"sort"
	"math/rand"
)

//声明Hero结构体
type Hero struct{
	Name string
	Age int
}
//声明一个Hero结构体切片类型
type HeroSlice []Hero


//实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

//按Ｈero的年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age > hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}


func main() {
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name:fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	fmt.Println(heroes)
	sort.Sort(heroes)
	fmt.Println(heroes)
}