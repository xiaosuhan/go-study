package main

import "fmt"

//定义一个底层类型是interface的 USB 类型
// USB 类型包括 Name()和 Connect()方法
type USB interface {
	Name() string
	Connect()
}

//定义一个底层数据类型是 struct的 PhoneConecter类型。这个类型里面有 name字段
type PhoneConecter struct {
	name string
}

//给PhoneConecter类型绑定一个 Name() string 的方法
func (pc PhoneConecter) Name() string  {
	return pc.name
}
//给PhoneConecter类型绑定一个Connect()的方法
func (pc PhoneConecter) Connect()  {
	fmt.Println("Connected: ",pc.name)
}

func main()  {
	//第一种显示的声明，实现了哪个接口。
	var a USB
	a = PhoneConecter{"PhoneConecter-name-a"}
	a.Connect()

	//第二种，不需要显示的声明。只要实现了接口的方法，系统会自动找到你是对应的哪个接口
	b := PhoneConecter{"PhoneConecter-name-b"}
	b.Connect()
	//调用测试传进去的是不是 USB类型的数据
	Disconnect(b)

}

func Disconnect(usb USB)  {
	fmt.Println("测试传入进来的是不是USB类型的")
}
