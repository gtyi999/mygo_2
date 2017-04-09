package main

import (
"fmt"

)

func main(){
	//Select
	//1--可以处理一个或多个channel 的发送与接收
	//2--同时有多个可用的channel时按随机顺序处理
	//3--可用空的select来阻塞main函数
	//4--可设置超时
	c1:=make(chan int)

	go func() {
		for v:=range c1{
			fmt.Println(v)
		}

	}()

	for{
		select {
		case c1<-0:
		case c<-1:
		}
	}

}
