package main

import "fmt"

func main()  {
	a := 1
	//定义一个 int型指针，& 是取地址符号。赋值给 p
	var p *int =&a

	//打印 p 和 *p
	fmt.Println(p)
	fmt.Println(*p)
}
