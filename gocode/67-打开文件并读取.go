package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)



func main() {
	//file的叫法
	//file对象, file指针, file文件句柄
	file, err := os.Open("./hello.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Printf("file=%v\n", file)
	
	defer file.Close() //要及时关闭file句柄,否则会有内存泄露
	//创建一个*Reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')//读到一个换行结束
		if err == io.EOF {//表示读到末尾了
			break
		}
		fmt.Print(str)
	}
}