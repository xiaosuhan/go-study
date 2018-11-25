package main

import "fmt"

func main(){
	a := 1
	switch {
	case a >=0 :
		fmt.Println("a=0")
		//加上这个表示，这个符合了，还会往下执行
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")

	}
}


