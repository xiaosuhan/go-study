package main

import "fmt"

type person struct {
	Name string
	Age int
	Contact struct{
		Phone, City string
	}
}

func main()  {
	a := person{Name: "suhan",Age: 16}
	a.Contact.Phone = "12312323"
	a.Contact.City = "shanghai"
	fmt.Println(a)
}

