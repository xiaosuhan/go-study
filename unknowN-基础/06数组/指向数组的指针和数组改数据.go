package main

import "fmt"

func main()  {
	//数组
	a := [5]int{}
	a[1] = 2
	fmt.Println(a)
	//指向数组的指针  两个一样都是可以改数据
	p := new([5]int)
	p[1] = 2
	fmt.Println(p)

}