package main

import "fmt"

func main()  {
	//空接口
	var a interface{}
	fmt.Println(a == nil)

	//执行nil 的指针，其实是有内存地址的
	var p *int = nil
	a = p
	fmt.Println(a == nil)
}