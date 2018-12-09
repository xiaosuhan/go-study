package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {
	//获取当前核数，多核运行
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 创建一个空的WaitGroup
	wg := sync.WaitGroup{}
	//给它定10个任务
	wg.Add(10)

	for i := 0;i < 10;i++ {
		//传WaitGroup 指针类型过去
		go Go3(&wg,i)
	}
	//等待都执行完
	wg.Wait()

}

//传进去WaitGroup 指针类型。
func Go3(wg *sync.WaitGroup,index int)  {
	a := 1
	for i := 0;i <1000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	// 标记这个任务做完了，就可以去掉了。
	wg.Done()
}
