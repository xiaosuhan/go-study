package main

import "fmt"

func main()  {
	// if 循环中定义的变量，作用范围只是在 if 循环中。超出了就不生效了
	if a := 1; a > 0{
		fmt.Println(a)
	}
	//打印的时候提示 a 未定义
	fmt.Println(a)
}

