package main

import "fmt"

func main()  {
	a := &struct {
		Name string
		Age int
	}{
		Name: "suhan",
		Age:25,
	}
	fmt.Println(a)
}

