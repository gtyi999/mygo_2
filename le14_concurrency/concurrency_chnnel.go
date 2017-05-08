package main

import (
	"fmt"
	"time"
)

var c chan int
func ready(w string,sec int){
	time.Sleep(time.Duration(sec*1e9))
	fmt.Println(w,"is ready!")
	c <- 1
}

var a2 string
var c2=make(chan int,10)
func f(){
	a2="hello,world"
	c2<-0
}

var a3 string
var c3=make(chan int)
func f3(){
	a3="hello,world"
	<-c3
}
func main(){
	fmt.Println("1--channel start")
	c=make(chan int)
	go ready("Tee",2)
	go ready("Coffee",1)

	fmt.Println("I m waiting,but not too long")
	<-c
	<-c
	fmt.Println("1--channel end")

	fmt.Println("2--有缓冲的channel start")
	go f()
	<-c2
	print(a2)
	fmt.Println("2--有缓冲的channel end")

	fmt.Println("3--无缓冲的channel start")
	//由于c3是无缓冲的channel，因此必须保证取操作<-c3 先于放操作c3<- 0
        go f3()
	c3<-0
	print(a3)
	fmt.Println("3--无缓冲的channel end")
}
