package main

import (
	"fmt"
	"runtime"
)

//说明：根据最后一次循环来判断程序是否执行完。
//结果：在单线程时没有问题，在多核上运行的时候就会出现问题

func main()  {
	//使用多核来运行。runtime.NumCPU() 是获取的当前机器的CPU核数
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := make(chan bool)
	for i := 0; i< 10 ;i++  {
		go Go2(c, i)
	}
	<- c
}

func Go2(c chan bool,index int)  {
	a := 1
	for i := 0;i <1100000000 ;i++  {
		a += i
	}
	fmt.Println(index,a)

	if index == 9 {
		c <- true
	}
}

