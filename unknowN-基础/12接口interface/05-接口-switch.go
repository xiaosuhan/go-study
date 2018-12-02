package main

import "fmt"

//定义一个底层类型是interface的 USB_interface 类型
// USB_interface 类型包括 Name_func()和 Connect_interface .用到了接口的嵌入
type USB_interface2 interface {
	Name_func2() string
	//把下面的接口，嵌入到USB_interface这个接口中
	Connect_interface2
}

//再定义一个声明Connect_interface的接口.里面有Connect_func 方法
type Connect_interface2 interface {
	Connect_func2()
} 

//定义一个底层数据类型是 struct的 PhoneConecter_struct 类型。这个类型里面有 name字段
type PhoneConecter_struct2 struct {
	name string
}

//给PhoneConecter_struct类型绑定一个 Name_func() string 的方法
func (pc PhoneConecter_struct2) Name_func2() string  {
	return pc.name
}
//给PhoneConecter_struct类型绑定一个Connect_func()的方法
func (pc PhoneConecter_struct2) Connect_func2()  {
	fmt.Println("Connected: ",pc.name)
}

func main()  {
	//第二种，不需要显示的声明。只要实现了接口的方法，系统会自动找到你是对应的哪个接口
	b := PhoneConecter_struct2{"PhoneConecter-name-b"}
	//b.Connect_func2()
	//调用测试传进去的是不是 USB类型的数据
	c := 1
	Disconnect_func2(b)
	Disconnect_func2(c)

}

//interface{} 空接口。 任何类型都实现了空接口，所以空接口=任何类型，可以接收任何类型
func Disconnect_func2(usb interface{})  {
	switch v := usb.(type)  {
	case PhoneConecter_struct2:
		fmt.Println("PhoneConecter_struct2--:", v.name)
	default:
		fmt.Println("未知的设备")
	}
	
}

