package  main

import (
	"fmt"

)

func main(){
	//Select
	//1--可以处理一个或多个channel 的发送与接收
	//2--同时有多个可用的channel时按随机顺序处理
	//3--可用空的select来阻塞main函数
	//4--可设置超时
	c1,c2:=make(chan int),make(chan string)
	o:=make(chan bool)
	go func(){
		for{
			select {
			case v,ok:=<-c1:
				if !ok{
					o<-true
					break
				}
				fmt.Println("c1",v)
			case v,ok:=<-c2:
				if !ok{
					o<-true
					break
				}
				fmt.Println("c2",v)
			}
		}
	}()

	c1<-1
	c2<-"hi"
	c1<-2
	c2<-"hello"

	close(c1)
	close(c2)
	<-o
}