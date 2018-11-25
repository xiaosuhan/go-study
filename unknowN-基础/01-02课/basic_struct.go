//当前程序的包名
package main

//导入其他的包，导入的包如果没有用到会报错。goland编译器会自动把不用的去掉
//用括号的是可以导入多个包，suhan 是起的别名。 省略调用 和别名不可以同时使用
import (
	suhan "fmt"
	)

// 省略调用--> 别名起成. 不方便看代码，不建议
//import . fmt


//常量的定义
const  PI  = 3.14

//全局变量的声明与赋值
var name = "gopher"

//一般类型声明,还可以自己声明一个类型？ 首字母小写表示私有的，只能同一个包内访问
type newType int

//结构体声明
type gopher struct {} 

//接口的声明
type goland interface {}

//由main函数作为程序入口点启动
func main() {
	suhan.Println("Hello World!你好，世界")
}
