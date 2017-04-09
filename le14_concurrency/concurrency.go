package main

import (
	"fmt"
)

func main(){
	/*
	//Channel
	//---Channel 是goroutine沟通的桥梁，大多数阻塞同步的
	//---通过make创建,close关闭
	fmt.Println("1---Channel  start")
	c:=make(chan bool)
	go func(){
		fmt.Println("Go Go Go!!!")
		c <- true
	}()
	<-c
	//---Channel是引用类型
	//---可以使用for range来迭代不断操作channel
	fmt.Println("2---for range  start")
	c1:=make(chan bool)
	go func(){
		fmt.Println("DO DO DO")
		c1<-true
		close(c1)
	}()
	for v :=range c1{
		fmt.Println(v)
	}
	fmt.Println("2---for range  end")

	//---可以设置单向或者双向通道

	//--可以设置缓存大小,在未被填满前不会发生阻塞
	//有缓存是异步的，没有缓存是同步的
	fmt.Println("3---设置chan 缓存 start")
	//c3:=make(chan bool,1)
	c3:=make(chan bool)
	go func(){
		fmt.Println("GO3 GO3")
		//c3<-true
		<-c3
	}()
	c3<-true
	fmt.Println("3--设置chan 缓存 end")*/

	fmt.Println("4---start ")
	c4:=make(chan  bool)
	for i:=0;i<10;i++{
		go Go(c4,i)
	}
	<-c4

}

func Go(c4 chan bool,index int){
	a:=1
	for i:=0;i<1000000;i++{
		a+=i
	}
	fmt.Println(index,a)
	if index==9{
		c4<-true
	}
}

