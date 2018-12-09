package main

import (
	"fmt"
	"time"
)

func main()  {
	c := make(chan bool)
	select {
	case v := <- c:
		fmt.Println(v)
	// 设置 select超时时间
	case <- time.After(3 * time.Second):
		fmt.Println("timeout")
		
	}
}
