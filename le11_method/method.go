package main

import (
	"fmt"
)

type A struct {
	B
	Name string
}
type  B struct {
	Name string
}
type  C struct {
	Name string
}

type D struct {
	B
	C
}

type m1 struct {
	Name string
}
type m2 struct {
	Name string
}

type TZ int

type AAAA struct {
	name string//小写为私有权限
}
func main(){
	fmt.Println("1---课堂作业 START")
        a:=A{Name:"A",B:B{Name:"B"}}
	fmt.Println(a.Name,a.B.Name)

	b:=D{B:B{Name:"B"},C:C{Name:"C"}}
	fmt.Println(b.B.Name,b.C.Name)
	fmt.Println("method")
	fmt.Println("1---课堂作业 end")

	fmt.Println("2---method start")
	mm1:=m1{}
	mm1.Print()
	mm2:=m2{}
	mm2.Print()
	fmt.Println("2---method end")
	fmt.Println("3----method 指针传递 start")
	mm3:=m1{}
	mm3.DoPrint()
	fmt.Println(mm3.Name)//改变Name的值
	fmt.Println(mm2.Name)//没有改变Name
	fmt.Println("3----method 指针传递 end")
	fmt.Println("4--类型别名 start")
	var cc TZ
	cc.Print()
	(*TZ).Print(&cc)
	fmt.Println("4--类型别名 end")
	fmt.Println("5--方法的访问权限  start")
	aaa:=AAAA{}
	aaa.Print()
	fmt.Println(aaa.name)
}

func (method1 m1)Print(){
	fmt.Println("method1")
}

func (method2 m2)Print(){
	method2.Name="method2"
	fmt.Println("method2")
}

func (method3 *m1)DoPrint(){
	method3.Name="method3"
}

func (a *TZ)Print(){
	fmt.Println("TZ")
}

func (aaa *AAAA)Print(){
	aaa.name="123"
	fmt.Println(aaa.name)
}