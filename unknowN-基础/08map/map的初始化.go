package main

import "fmt"

func main()  {
	m := make(map[int]string)
	m[1] = "ok"
	a := m[1]
	fmt.Println(a)
	fmt.Println(m)

	delete(m,1)
	// 删除以后再打印 value就是空
	b := m[1]
	fmt.Println(b)
	fmt.Println(m)
}