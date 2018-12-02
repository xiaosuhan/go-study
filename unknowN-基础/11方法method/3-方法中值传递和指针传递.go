package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
} 

func main()  {
	a := A{}
	a.Suhanprint()
	fmt.Println(a.Name)

	b := B{}
	b.Suhanprint()
	//这个打印出来是空的，说明在方法中值传递，不会修改原始的值，出了方法就没用了
	fmt.Println(b.Name)

}

//把 a 和 A传进来，就把Suhanprint 和 a绑定到一起了
//传进来的东西不同，就是不同的传进来的东西的方法，这样就算方法名称一样，也不冲突。
func (a *A) Suhanprint()  {
	a.Name = "指针的Name"
	fmt.Println("增加的方法")
}

func (b B) Suhanprint()  {
	b.Name = "值的Name"
	fmt.Println("增加的方法第二个")
}



