package main
import (
	"fmt"
)


func main(){
	for i := 1; i <= 9; i++ {
		for z := 1; z <= i; z++ {
			fmt.Printf("%d * %d = %d\t",i ,z, i * z)
		}
		fmt.Println("")
	}
}