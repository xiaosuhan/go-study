package main

import "fmt"

type TZ int

func main()  {
	var a TZ
	a.Suhanprint()
	(*TZ).Suhanprint(&a)
}

func (a TZ) Suhanprint()  {
	fmt.Println("int类型增加的方法")
}
