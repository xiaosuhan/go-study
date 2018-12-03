package main

import "fmt"

func main()  {
	//第一步是把之前的函数直接拿过来，名字去掉
	a := func() {
		fmt.Println("Func 匿名")
	}
	a()

	b := func(){
		fmt.Println("Func 匿名11")
	}
	b()

}
