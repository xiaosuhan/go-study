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
	//这里传入的是个值，不是个指针，如果是个指针的话，Info里面就不能那样写了。
	// 所以可以加个类型的判断
	Info(u)
}

func Info(g interface{})  {
	t := reflect.TypeOf(g)
	fmt.Println("Type:", t.Name())

	//如果传入的不是值，是指针了，需要加个判断
	if k := t.Kind();k != reflect.Struct {
		fmt.Println("类型不对")
		return
	}
	
	v := reflect.ValueOf(g)
	fmt.Println("Fields:-----")

	for i := 0;i < t.NumField(); i++{
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0;i < t.NumMethod(); i++  {
		m := t.Method(i)
		fmt.Printf("%6s : %v\n", m.Name,m.Type)
	}



}