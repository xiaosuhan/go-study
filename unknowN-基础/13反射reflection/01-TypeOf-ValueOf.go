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

//给User类型绑定一个方法
func (u User) Hello()  {
	fmt.Println("hello 绑定的方法")
}

func main()  {
	//没有写key的话，就是按照位置来赋值的。
	u := User{1, "suhan", 13}
	Info(u)
}

func Info(g interface{})  {
	t := reflect.TypeOf(g)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(g)
	fmt.Println("Fields:-----")

	for i := 0;i < t.NumField(); i++{
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}





}