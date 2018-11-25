package main

import "fmt"

//方法体外全局变量 var 关键字必不可少
// 冒号就是替代 var 关键字的，这里有var 关键字了就不能使用冒号了
var (
	f,g,h = 2,3,4
)

func main(){
	//多个变量并行声明与赋值。
	// 方法体内可以使用冒号替代var。 方法体外全局变量 var 关键字必不可少
	// _ 是把对应的值忽略
	a,_,c,d := 1,2,3,4
	fmt.Println(a)

	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}