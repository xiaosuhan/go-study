package main

import "fmt"

func main() {
	var a int	//变量的声明
	a = 123 //变量的赋值
	fmt.Println(a)

	//声明和赋值同时
	var b int = 21
	fmt.Println(b)

	//省略类型，让系统推断
	var e  = 21
	fmt.Println(e)

	//声明和赋值最简单的写法
	c := 23
	fmt.Println(c)
}

