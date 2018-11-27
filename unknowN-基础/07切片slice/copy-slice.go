package main

import "fmt"

func main()  {
	s1 := []int{1,2,3,4}
	s2 := []int{5,6,7,8,9,0,12,3}
	copy(s1,s2)
	//可以指定拷贝多少位
	copy(s1[2:4],s2[1:3])
	fmt.Println(s1)
}