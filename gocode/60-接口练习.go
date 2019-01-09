package main
import (
	"fmt"
	"math/rand"
	"sort"
)


//定义一个整数的切片类型
type intSlice []int

// type Interface interface {
//     // Len方法返回集合中的元素个数
//     Len() int
//     // Less方法报告索引i的元素是否比索引j的元素小
//     Less(i, j int) bool
//     // Swap方法交换索引i和j的两个元素
//     Swap(i, j int)
// }

func (intslice intSlice) Len() int {
	return len(intslice)
}
func (intslice intSlice) Less(i, j int) bool {
	return intslice[i] > intslice[j]
}
func (intslice intSlice) Swap(i, j int) {
	// intslice[i] = intslice[i] + intslice[j]
	// intslice[j] = intslice[i] - intslice[j]
	// intslice[i] = intslice[i] - intslice[j]
	intslice[i], intslice[j] = intslice[j], intslice[i]
}

func main() {
	var list intSlice
	for i := 0; i < 10; i++ {
		list = append(list, rand.Intn(20))
	}
	fmt.Println(list)
	sort.Sort(list)
	fmt.Println(list)
}