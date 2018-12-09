package main

import (
	"fmt"
	"reflect"
)

//每个不懂的值都打印一下，看看是个什么东西

//定义一个struct类型的变量User
type User1 struct {
	Id int
	Name string
	Age int
}

func main()  {
	//声明和赋值
	u := User1{1,"ok",12}
	Set(&u)
	fmt.Println(u)
}

//形参是一个空接口，空接口就是说什么类型都可以接收
func Set(o interface{}) {
	v := reflect.ValueOf(o)
	//main.User 是类型
	fmt.Println("reflect.TypeOf(o):如下")
	fmt.Println(reflect.TypeOf(o))
	//{1 ok 12}  是值
	fmt.Println("reflect.ValueOf(o):如下")
	fmt.Println(v)
	//传进来的是值，kind 就是 struct，传进来的是指针 kind就是 ptr。 ptr表示指针的意思
	fmt.Println("v.Kind():如下")
	fmt.Println(v.Kind())

	fmt.Println("reflect.Ptr：如下")
	fmt.Println(reflect.Ptr)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet(){
		fmt.Println("传进来是不是指针，或者传进来的是指针，但是不能被修改值")
	} else {
		//v 就是{1 ok 12}。 如果v 是指针。v.Elem()就表示其指向的元素，就是所有的值
		v = v.Elem()
		fmt.Println("v.Elem()：如下")
		fmt.Println(v)
	}

	fmt.Println("在if循环外面，原来是&{1 ok 12}，通过v.Elem()变成了{1 ok 12}：如下")
	fmt.Println(v)

	//这是两行合在一起写了。v就是 User的struct里面的数据。 FieldByName 就是表示根据key来过滤值
	fmt.Println("FieldByName:id 和 Age的值")
	fmt.Println(v.FieldByName("Id"))
	fmt.Println(v.FieldByName("Age"))


	f := v.FieldByName("Name1")
	if !f.IsValid(){
		fmt.Println("Name1 的key对应的值不存在")
		return
	}

	if f.Kind() == reflect.String{
		f.SetString("BYEBYE")
	}



}

