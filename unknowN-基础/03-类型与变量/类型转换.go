package main

import "fmt"

func main(){
	//  :=  之前没有声明过，可以使用 冒号，之前声明过了就不需要冒号了。

	// float 转 int 是可以的，但是 小数点后面的就没有了
	var a float32 = 1.1
	b := int(a)
	fmt.Println(b)

	// bool 转成int 是不行的
	var c bool = true
	d := int(c)
	fmt.Println(d)

}
