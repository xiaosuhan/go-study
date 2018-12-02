package main

import "fmt"

//定义一个SH 类型，底层数据类型是 int
type SH int

//sh *SH 是把这个方法和SH 类型绑定。 * 表示指针传递。 sh 是谁在调用。
//num int 是函数的形参
func (sh *SH) Increase(num int)  {
	//虽然 SH 类型和 int 底层都是一样的，但是还是要转一下，是两个不同的东西
	// sh 就是a，num 是100
	*sh += SH(num)
}

func main()  {
	var a SH
	// sh 就是a，num 是100
	a.Increase(100)
	fmt.Println(a)
}

