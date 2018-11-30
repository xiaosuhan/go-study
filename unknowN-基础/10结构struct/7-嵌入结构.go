package main

import "fmt"

//公共的
type human struct {
	Sex string
}

type student struct {
	//把公共的嵌入到这个里面。类似于继承
	human
	Name string
	Age int
}

type teacher struct {
	//把公共的嵌入到这个里面。类似于继承
	human
	Name string
	Age int
}

func main()  {
	a := student{Name: "suhan",Age: 25,human: human{Sex: "man"}}
	b := teacher{Name: "suhan2", Age: 30,human: human{Sex: "man"}}
	fmt.Println(a)
	fmt.Println(b)

	//可以通过以下方式来更改
	a.Name = "qqq"
	a.Age = 11
	a.human.Sex = "woman"
	a.Sex = "1111"
	fmt.Println(a)
	
}