package main

import "fmt"

func main()  {
	c3 := make(chan int)
	go func() {
		for v := range c3{
			fmt.Println(v)
		}
	}()

	//select是执行一次就结束的，所以要加个for循环
	for {
		select {
		case c3 <- 0:
		case c3 <- 1:
		}
	}

}
