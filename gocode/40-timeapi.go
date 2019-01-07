package main
import (
	"fmt"
	"time"
)


func main () {
	now := time.Now()
	// fmt.Println(time1)
	fmt.Printf("%v %T\n", now, now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	weekday := now.Weekday()
	fmt.Println(year, month, day, weekday)
	fmt.Printf("%d-%d-%d\n", now.Year(), now.Month(), now.Day())
	fmt.Println(now.Format("2006/01/02 15:04:05"), "\n", now.Unix())
}