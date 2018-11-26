package main

import "fmt"

func main()  {
	//多维数组
	a := [2][3]int{
		{1:1},
		// }} 一定要贴在一起
		{2:1}}

	fmt.Println(a)
}
