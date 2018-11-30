package main

import "fmt"

func main()  {
	//for i := 0;i < 3;i++ {
	//	defer fmt.Println(i)
	//}
	//三次打印出来的都是 3
	for j := 0;j < 3;j++ {
		defer func() {
			fmt.Println(j)
		}()
	}

}

