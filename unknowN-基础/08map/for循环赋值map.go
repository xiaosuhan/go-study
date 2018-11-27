package main

import "fmt"

func main()  {
	sm := make([]map[int]string,5)
	// 只有操作 k 才能真实的操作 map的值
	for k := range sm{
		sm[k] = make(map[int]string)
		sm[k][1] = "ok"
		fmt.Println(sm[k])
	}
	fmt.Println(sm)
}
