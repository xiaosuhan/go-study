package main

import "fmt"

func main()  {

}
//第一个括号里面是传入的参数，第二个括号里面的是 返回值
func A(a,b int) (c,d int)  {
	fmt.Println("11")
}

//第一个括号里面是传入的参数，第二个int是返回值，表示返回一个int的值
func B(a,b int) int  {
	fmt.Println("11")
}

//第一个括号里面是传入的参数，第二个括号里面的是 返回值
func C(a,b int) (c,d int)   {
	//因为返回值里面写了 c d 表示已经存在了，就不用在声明了，就不用 冒号了 直接赋值
	c, d = 1, 2
	//返回的时候为了增加可读性，还是把返回值都写上
	return c,d
}


