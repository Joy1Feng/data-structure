package main
import (
	"fmt"
	"os"
)


func main() {
	//file的叫法
	//file对象, file指针, file文件句柄
	file, err := os.Open("./hello.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Printf("file=%v\n", file)
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}