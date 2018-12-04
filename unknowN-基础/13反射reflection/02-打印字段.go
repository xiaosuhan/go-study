package main

import (
	"fmt"
	"reflect"
)

type User2 struct {
	Id int
	Name string
	Age int
}

type Manager struct {
	User2
	title string
}

func main()  {
	m := Manager{User2:User2{1,"suhan",26}, title: "123"}
	t := reflect.TypeOf(m)
	//t 是 main.Manager
	fmt.Println(t)

	//打印Manager 里面的两个字段。
	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.Field(1))
	fmt.Println("-----------")
	//打印 Manager中的User2中的 Id
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0,0}))
	//打印 Manager中的User2中的 Name
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0,1}))


}


