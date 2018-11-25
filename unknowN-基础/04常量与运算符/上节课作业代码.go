package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var a int = 65
	// int --> string
	b  := strconv.Itoa(a)

	// string --> int
	a,_ = strconv.Atoi(b)
	fmt.Println(a)
}
