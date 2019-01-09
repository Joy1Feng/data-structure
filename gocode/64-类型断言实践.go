package main
import (
	"fmt"
)


//定义了一个接口
type Usb interface{
	//生声明了两个没有实现的方法
	Start()
	Stop()
}

type Usb2 interface{
	//生声明了两个没有实现的方法
	Start()
	Stop()
}

//让phon实现Usb接口方法
type Phone struct{

}
func (phone Phone) Start() {
	fmt.Println("手机开始工作")
}
func (phone Phone) Stop() {
	fmt.Println("手机停止工作")
}
func (phone Phone) Call() {
	fmt.Println("打电话...")
}

//让camera实现Usb接口方法
type Camera struct{

}
func (camera Camera) Start() {
	fmt.Println("相机开始工作")
}
func (camera Camera) Stop() {
	fmt.Println("相机停止工作")
}

//计算机
type Computer struct{

}
//编写一个working,接受一个Usb接口类型变量
//只要是实现了usb接口(所谓实现了usb接口,就是实现了usb接口声明所有方法)
func (c Computer) working(usb Usb) {
	//通过usb接口变量来调用start和stop方法
	usb.Start()
	//如果usb是执行phone结构体变量, 则还需要调用call方法
	phone, ok := usb.(Phone)
	if ok {
		phone.Call()
	}
	usb.Stop()
}


func main() {
	//测试
	//创建结构体实例
	computer := Computer{}
	phone := Phone{}
	computer.working(phone)
	camera := Camera{}
	computer.working(camera)
}