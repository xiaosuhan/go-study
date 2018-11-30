package main

import "fmt"

type person1 struct {
	Name string
	Age int
}


func main()  {
	a := person1{Name: "jobe", Age: 19}
	b := person1{Name: "jobe", Age: 14}
	fmt.Println(a == b)
}


