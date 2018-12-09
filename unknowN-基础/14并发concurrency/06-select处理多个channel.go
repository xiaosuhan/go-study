package main

import (
	"fmt"
)

func main()  {
	c1, c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		// 死循环，不断的去执行
		for {
			// select处理多个channel
			select {
			//取 c1 channel中的值
			case v, ret := <- c1:
				if !ret {
					//channel关闭了报错就会 给 channel o中存一个值。
					o <- true
					break
				}
				fmt.Println("c1---", v)

			case v, ret := <- c2:
				if !ret {
					//channel关闭了报错就会 给 channel o中存一个值。
					o <- true
					break
				}
				fmt.Println("c2---",v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 13
	c2 <- "hello"
	//可能出现，channel里面放进去值，上面的代码还没有执行完，就关闭了。
	close(c1)
	close(c2)

	// 读取到值。就说明 c1 c2中有一个channel关闭了，然后就程序退出。
	// case中不成功的时候会break，所以也会程序结束。
	<- o

	//没有 cahnnel o  和 break的话，程序就不会退出了

}