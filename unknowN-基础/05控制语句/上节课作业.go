package main

import "fmt"

// iota 实现计算机存储单位的枚举

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
)

func main()  {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
}

