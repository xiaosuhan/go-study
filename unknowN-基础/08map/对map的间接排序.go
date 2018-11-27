package main

import (
	"fmt"
	"sort"
)

func main()  {
	//创建并初始化一个 map
	m := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
	//初始化一个 slice
	s := make([]int,len(m))
	i := 0
	for k,_:= range m{
		s[i] = k
		i++
	}
	//给 slice 排序
	sort.Ints(s)
	fmt.Println(s)



}