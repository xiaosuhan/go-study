package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello(name string)  {
	fmt.Println("Hello",name,",my name is ",u.Name)
}

func main()  {
	u := User{1,"ok",12}
	fmt.Println("正常调用一个方法的结果：")
	u.Hello("Joe")

	v := reflect.ValueOf(u)
	fmt.Println("------")
	fmt.Println(v)
	mv := v.MethodByName("Hello")
	fmt.Println(mv)
	fmt.Println(reflect.ValueOf("joe"))
	fmt.Println("======")

	args := []reflect.Value{reflect.ValueOf("joe")}
	fmt.Println(args)
	mv.Call(args)


}







