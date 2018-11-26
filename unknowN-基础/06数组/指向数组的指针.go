package main

import "fmt"

func main(){
	a := [...]int{5:1}
	var p *[6]int = &a
	fmt.Println(p)
}

