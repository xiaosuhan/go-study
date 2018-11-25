package main

import "fmt"

func main()  {
	a := 1
	for {
		a++
		if a > 3{
			break
		}
		fmt.Println(a)
	}
	fmt.Println("over")
}