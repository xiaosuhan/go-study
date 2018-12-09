package main

import (
	"fmt"
	"time"
)

func main()  {
	//goroutine 的使用
	go Go()

	//暂停两秒.因为main执行太快了，没有等 Go函数执行完，所以如果不加这个，什么都不会输出
	time.Sleep(2 * time.Second)
}

func Go()  {
	fmt.Println("go go go")
}
