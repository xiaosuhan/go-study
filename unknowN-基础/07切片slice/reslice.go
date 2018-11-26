package main

import "fmt"

func main()  {
	a :=[]byte{'a','b','c','d','e','f','g','h','i','f'}
	sa := a[2:5]
	fmt.Println(string(sa))
	fmt.Println(len(sa),cap(sa))
	//索引3 超过了 sa 切片的个数。但是没有超出容量，所以是可以继续读取到数的。
	sa2 := sa[3:4]
	fmt.Println(string(sa2))

}
