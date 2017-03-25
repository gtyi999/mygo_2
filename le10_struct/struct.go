package main

import (
	"fmt"
	//"github.com/syndtr/goleveldb/leveldb/testutil"
)

type test struct {

}

type person struct {
	Name string
	Age  int
}

type person2 struct {
	Name string
	Age int
	Contact struct{
		Phone,City string
		}
}

type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age int
}
type student struct {
	human
	Name string
	Age int
}

func main(){
	fmt.Println("1---struct 初始化 start")
	a:=test{}
	fmt.Println(a)
	b:=person{}//string类型的默认值为空  int的默认值为0
	fmt.Println(b)
	c:=person{"john",19}
	fmt.Println(c)
	c.Name="micle"
	c.Age=20
	fmt.Println(c)
	fmt.Println("1---struct 初始化 end")
	fmt.Println("2---struct 值拷贝 start")
	d:=person{
		Name:"joe",
		Age:19,
	}
	fmt.Println(d)
	A(d)
	fmt.Println(d)
	fmt.Println("2---struct 值拷贝 end")
	fmt.Println("3--指针传递，真正的修改 start")
	B(&d)
	fmt.Println(d)
	fmt.Println("3--指针传递，真正的修改 start")

	fmt.Println("4---初始化struct的地址 start")
	e:=&person{
		Name:"bbb",
		Age:12,
	}
	e.Name="ok"
	fmt.Println(e)
	B(e)
	fmt.Println(e)
	fmt.Println("4---初始化struct的地址 end")
	fmt.Println("5--匿名结构 ---start")
	f:= &struct {
		Name string
		Age int
	}{
		Name:"joe",
		Age:19,
	}
	fmt.Println(f)
	fmt.Println("5--匿名结构 ---end")
	fmt.Println("6--嵌套结构--start")
	g:=person2{Name:"fff",Age:19}
	g.Contact.Phone="1234567"
	g.Contact.City="beijing"
	fmt.Println(g)
	fmt.Println("6--嵌套结构--end")
	fmt.Println("7---结构的比较 start")
	h:=person{Name:"joe",Age:19}
	I:=person{Name:"joe",Age:20}
	fmt.Println(h==I)
	fmt.Println("7---结构的比较 end")
	j:=teacher{Name:"joe",Age:19,human:human{Sex:1}}
	k:=student{Name:"joe",Age:20,human:human{Sex:0}}
	fmt.Println(j,k)
	j.Name="joe2"
	j.Age=13
	j.human.Sex=100
	fmt.Println(j)


}

//值拷贝函数
func A(per person){
	per.Age=13
	fmt.Println("A",per)
}
func B(per *person)  {
	per.Age=13
}
