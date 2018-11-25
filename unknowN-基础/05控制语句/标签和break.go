package main

import "fmt"

func main(){
	//break是跳出LABEL1标签同一级的for循环。跟 goto不一样
	LABEL1:
		for {
			for i := 0;i < 10 ;i++  {
				if i > 3{
					break LABEL1
				}
			}
		}
	fmt.Println("over")
}
