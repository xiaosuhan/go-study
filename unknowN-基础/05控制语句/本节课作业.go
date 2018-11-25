package main

import "fmt"

func main(){
LABEL:
		for i := 0;i < 10 ;i++  {
			for  {
				fmt.Println(i)
				//goto和continue的区别 goto就是死循环打出来的都是0
				//把 LABEL放在 循环后面就不会死循环了
				goto LABEL
				//continue LABEL
			}
		}

}