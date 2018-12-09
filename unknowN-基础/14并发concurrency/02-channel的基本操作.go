
package main

import "fmt"

func main(){
	// 定义一个bool型的channel。就是定义一个通道，可以传入和取出在里面的值
	// 用的是 make 不用自己再手动关闭了，程序执行完就释放了。因为这里只执行一次，一般是需要close 关闭的
	c := make(chan bool)
	go func() {
		fmt.Println("go go go!!!!")
		// 给 c 这个channl 存入一个 true
		c <- true
	}()

	// 取这步，会阻塞。直到 执行了匿名函数中的 存的操作。才能取出来。
	//所以这里没有使用 sleep 也是可以的
	<- c

}



