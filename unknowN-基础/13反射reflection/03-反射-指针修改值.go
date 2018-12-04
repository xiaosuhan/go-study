package main

import (
	"fmt"
	"reflect"
)

func main()  {
	x := 123
	//反射获取值。有指针和非指针两种
	v1 := reflect.ValueOf(x)
	fmt.Println(v1)
	v2 := reflect.ValueOf(&x)
	fmt.Println(v2)
	fmt.Println("--------")

	// 这样就可以修改原来的值
	v2.Elem().SetInt(9999)
	fmt.Println(x)

}

