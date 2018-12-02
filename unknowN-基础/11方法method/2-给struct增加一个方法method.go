package main

import (
	"fmt"
)

type A struct {
	Name string
}

func main()  {
	a := A{}
	a.Suhanprint()
}

//把 a 和 A传进来，就把Suhanprint 和 a绑定到一起了
//传进来的东西不同，就是不同的传进来的东西的方法，这样就算方法名称一样，也不冲突。
func (a A) Suhanprint()  {
	fmt.Println("增加的方法")
}

