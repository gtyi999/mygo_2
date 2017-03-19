package main

import (
	"fmt"
)

func main(){
	fmt.Println("0--课堂作业 start")
	m1:=map[int]string{1:"a",2:"b",3:"c",4:"d"}
	fmt.Println(m1)
	m2:=make(map[string]int)

	for k,v:=range m1{
		m2[v]=k
	}
	fmt.Println(m2)
	fmt.Println("0--课堂作业 end")
	fmt.Println("1--function  start")
	a,b:=1,2
	D(a,b)
	fmt.Println(a,b)
	fmt.Println("1--function end")
	fmt.Println("2--function slice值传递 start")
	s1:=[]int{1,2,3,4}
	fmt.Println(s1)
	E(s1)
	fmt.Println(s1)
	fmt.Println("2--function slice值传递 end")

}

func A(a int,b string) int {
	//返回值
	return 1
}
func B(a int,b string) (e,f,g int) {
	//e,f,g:=1,2,3//这不能作定义赋值
	e,f,g=1,2,3
	return e,f,g
}
func C(b string,a ...int){
	fmt.Println(a)
}
func D(s ...int){
	s[0]=3
	s[1]=4
	fmt.Println(s)
}
func E(s []int){
	s[0]=5
	s[1]=6
	s[2]=7
	s[3]=8
}

