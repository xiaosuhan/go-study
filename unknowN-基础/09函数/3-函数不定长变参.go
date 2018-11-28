package main

import "fmt"

func main()  {
	A("suhan",1,2,34)
}

//... 表示不定长变参。同 slice中的[...]相似。 同python中的kwargs相似
func A(b string, a ...int)  {
	fmt.Println(a)
}