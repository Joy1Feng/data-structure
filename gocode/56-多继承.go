package main
import (
	"fmt"
)

type Goods struct{
	Name string
	Price float64
}

type Brands struct{
	Name string
	Address string
}

type TV struct{
	Goods
	Brands
}

func main() {
	var tv TV
	tv.Goods.Name = "海尔"
	tv.Goods.Price = 9999.90
	tv.Brands.Name = "中国"
	tv.Brands.Address = "重庆"
	fmt.Println(tv.Goods.Name)
}