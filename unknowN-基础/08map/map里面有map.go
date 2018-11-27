package main

import "fmt"

func main()  {
	m := make(map[int]map[int]string)
	// 多返回值的用途
	a, ret := m[2][1]
	fmt.Println(a,ret)
	if !ret{
		m[2] = make(map[int]string)
	}
	m[2][1] = "GOOD	"
	a, ret = m[2][1]
	fmt.Println(a,ret)
}
