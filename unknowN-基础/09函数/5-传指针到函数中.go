package main

import "fmt"

func main()  {
	a := 1
	A(&a)
	fmt.Println(a)
}

//传递一个指针进来
func A(a *int)  {
	*a = 2
	fmt.Println(*a)
}
