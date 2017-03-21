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

	fmt.Println("3--function 指针传递 start")
	aa:=1
	G(&aa)
	fmt.Println(aa)
	fmt.Println("3--function 指针传递 end")

	fmt.Println("4--函数类型 start")
	funcH:=H
	funcH()
	fmt.Println("4--函数类型 end")

	fmt.Println("5--匿名函数 start")
	funcI:=func(){
		fmt.Println("func I")
	}
	funcI()
	fmt.Println("5--匿名函数 start")

	fmt.Println("6--闭包函数 start")
	funcJ:=closure(10)
	fmt.Println(funcJ(1))
	fmt.Println(funcJ(2))
	fmt.Println("6--闭包函数 end")

	fmt.Println("7--defer 关键字 start")
	//1-----
	//fmt.Println("a")
	//defer fmt.Println("b")
	//defer fmt.Println("c")
        //2---
	//for ii:=0;ii<3;ii++{
	//	defer fmt.Println(ii)
	//}
	//3--
	//for qq:=0;qq<3;qq++{
	//	defer func(){
	//		fmt.Println(qq)
	//	}()
	//}
	K()
	L()
	M()
	fmt.Println("7--defer 关键字 end")

	fmt.Println("8--课堂作业 start")
	var fs=[4]func(){}
	for jj:=0;jj<4;jj++{
		defer fmt.Println("defer jj= ",jj)
		defer func(){fmt.Println("defer_closure jj= ",jj)}()
		fs[jj]=func(){fmt.Println("closure jj= ",jj)}
	}
	for _,f:=range fs{
		f()
	}
	




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

func G(a *int){
	*a=2
	fmt.Println(a)
}
func H(){
	fmt.Println("func A")
}

func closure(x int) func(int) int{
	fmt.Println("%p\n",&x)
	return func(y int)int{
		fmt.Println("%p\n",&x)
		return x+y
	}
}

func K(){
	fmt.Println("Func K")

}
func L(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println("Recover in L")
		}
	}()
	panic("Panic in L")

}

func M(){
	fmt.Println("Func M")
}