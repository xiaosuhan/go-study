package main

import "fmt"


func main() {
	a := "string"
	// 不用直接把 len(a) 放在for循环中，循环次数多的话影响性能，因为每次都会执行一遍len
	L := len(a)
	for i := 0;i < L; i++{
		fmt.Println(a)
	}
	fmt.Println("over")
}


