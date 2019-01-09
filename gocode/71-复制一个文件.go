package main
import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strconv"
	"time"
	"io/ioutil"
)


func MyWrite(filepath string) {
	file, err := os.OpenFile(filepath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件错误,错误类型是:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("i am" + strconv.Itoa(i))
		writer.WriteString("\n")
	}
	writer.Flush()
	fmt.Println("写入内容成功")
}

func MyRead(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("打开文件错误,错误类型是:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("%v", data)
	}
}


func MyRead1(filepath string) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("打开文件失败错误类型是:", err)
	}
	// fmt.Printf("%s", []byte(content))
	fmt.Printf("%v", string(content))//将字节转换为字符串
}


func MyCopy(filepath1 string, filepath2 string) {
	file2, err2 := os.OpenFile(filepath2, os.O_WRONLY | os.O_APPEND, 0666)
	if err2 != nil {
		fmt.Println("打开文件错误,错误类型是:", err2)
		return
	}
	defer file2.Close()
	writer := bufio.NewWriter(file2)
	// for i := 0; i < 10; i++ {
	// 	writer.WriteString("i am" + strconv.Itoa(i))
	// 	writer.WriteString("\n")
	// }
	// writer.Flush()
	// fmt.Println("写入内容成功")

	file1, err1 := os.Open(filepath1)
	if err1 != nil {
		fmt.Println("打开文件错误,错误类型是:", err1)
		return
	}
	defer file1.Close()
	reader := bufio.NewReader(file1)
	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		time.Sleep(1 * time.Second)
		writer.WriteString(data)
	}
	writer.Flush()
	fmt.Println("写入内容成功")
}

func main() {
	// MyWrite("./joy.txt")
	// fmt.Println("....")
	// time.Sleep(2 * time.Second)
	// fmt.Println("+++")
	// MyRead("./writeread.txt")
	// time.Sleep(2 * time.Second)
	// MyRead1("./writeread.txt")
	// // fmt.Println(res)
	MyCopy("./joy.txt", "./joy1.txt")
}