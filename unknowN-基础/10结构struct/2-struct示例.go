package main

import "fmt"

type person struct {
	Name string
	Age int
}

func main()  {
	a := person{
		//初始化的时候赋值
		Name: "suhan",
	}
	//初始化以后赋值
	//a.Name = "suhan"
	a.Age = 26
	fmt.Println(a)
}
