package main

import "fmt"

func main()  {
	//容量为2 的空数组
	a := [2]int{}
	fmt.Println(a)
	//容量为5的数组，前两个数是1
	b := [5]int{1,1}
	fmt.Println(b)
	//索引3 就是第四个数组是设置为1，其他没有设置值就都是0
	c := [5]int{3:1}
	fmt.Println(c)
	//... 表示是长度由系统推断
	d := [...]int{5:1}
	fmt.Println(d)
}
