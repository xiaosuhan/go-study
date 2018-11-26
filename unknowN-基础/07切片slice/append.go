package main

import "fmt"

func main()  {
	s1 := make([]int,3,8)
	fmt.Printf("%p\n", s1)

	s1 = append(s1,1,2,3,)
	fmt.Printf("%v %p\n", s1, s1)

	s1 = append(s1,1,2,3,)
	fmt.Printf("%v %p\n", s1, s1)
}