package main
import (
	"fmt"
	"bufio"
	"os"
)


func main() {
	//创建一个新文件, 写入５句hello world
	file, err := os.OpenFile("./5hello.txt", os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件错误")
		return
	}
	defer file.Close()
	str := "hello, world"
	//写入时,使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
		writer.WriteString("\n")
	}
	//因为Writer是带缓存,因此在调用WriterString,其实内容先写入到缓存，
	//所以需要调用Flush方法,讲缓存的数据真正写入到文件中,否则文件中会丢失数据
	writer.Flush()
}
