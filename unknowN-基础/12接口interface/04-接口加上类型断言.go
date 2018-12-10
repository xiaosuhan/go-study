package main

import (
	"fmt"
	"reflect"
)

//定义一个底层类型是interface的 USB_interface 类型
// USB_interface 类型包括 Name_func()和 Connect_interface .用到了接口的嵌入
type USB_interface interface {
	Name_func() string
	//把下面的接口，嵌入到USB_interface这个接口中
	Connect_interface
}

//再定义一个声明Connect_interface的接口.里面有Connect_func 方法
type Connect_interface interface {
	Connect_func()
} 

//定义一个底层数据类型是 struct的 PhoneConecter_struct 类型。这个类型里面有 name字段
type PhoneConecter_struct struct {
	name string
}

//给PhoneConecter_struct类型绑定一个 Name_func() string 的方法
func (pc PhoneConecter_struct) Name_func() string  {
	return pc.name
}
//给PhoneConecter_struct类型绑定一个Connect_func()的方法
func (pc PhoneConecter_struct) Connect_func()  {
	fmt.Println("Connected: ",pc.name)
}

func main()  {
	//第二种，不需要显示的声明。只要实现了接口的方法，系统会自动找到你是对应的哪个接口
	b := PhoneConecter_struct{"PhoneConecter-name-b"}
	b.Connect_func()
	//调用测试传进去的是不是 USB类型的数据
	Disconnect_func(b)

}

func Disconnect_func(usb USB_interface)  {

	fmt.Println("=====")
	fmt.Println(usb)
	fmt.Println(reflect.TypeOf(usb))

	//如果传入的类型是对的。那ret 就是真，就执行括号里面的，直接return。不会执行最后一句打印的
	if pc, ret := usb.(PhoneConecter_struct);ret{
		fmt.Println(pc.name)
		fmt.Println(pc)
		return
	}

	//如果传入的类型不对，ret是false就不会执行上面括号里面的。就会下面这个打印的
	fmt.Println("传入的类型不对")
}

