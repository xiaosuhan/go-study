package main

import "fmt"

func main() {
	//有容量的叫数组
	var b [1]int
	fmt.Println(b)
	//没有容量的是切片 。零值也不一样
	var b2 []int
	fmt.Println(b2)
	var b3 [1]bool
	fmt.Println(b3)


}