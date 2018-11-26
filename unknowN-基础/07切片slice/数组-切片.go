package main

import "fmt"

func main()  {
	a := [10]int{1,2,3,4,5,6,7,8,9,0}
	fmt.Println(a)
	//索引为5-索引为9的。不包括索引10
	sli := a[5:10]
	fmt.Println(sli)
	//索引5到整个数组的尾部
	sli2 := a[5:]
	fmt.Println(sli2)
	//索引0到索引5，不包括索引5
	sli3 := a[:5]
	fmt.Println(sli3)

}
