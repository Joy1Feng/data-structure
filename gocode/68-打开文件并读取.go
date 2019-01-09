package main
import (
	"fmt"
	"io/ioutil"
)


func main() {
	//没有open和close都被封装到函数里面了
	file := "./hello.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("打开文件失败")
	}
	fmt.Printf("%s\n", []byte(content))//将字节转换为字符串
	fmt.Printf("%v", string(content))//将字节转换为字符串

}