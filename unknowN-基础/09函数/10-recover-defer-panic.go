package main

import "fmt"

//手动panic让程序报错，再通过defer和recover把程序恢复回来。

func main()  {
	A()
	B()
	C()
}

func A()  {
	fmt.Println("Func in A")
}
func B()  {
	// defer 定义一个匿名函数来recover
	defer func() {
		// 调用recover，如果err不为nil，就说明引发了panic。否则就不需要处理
		if err := recover();err != nil {
			fmt.Println("recover in B")
		}
	}()

	fmt.Println("xxxxxx")
	//手动触发一个panic
	//panic("panic in B")
}

func C()  {
	fmt.Println("Func in C")
}