package main

import (
	"fmt"
)

type test struct {

}

type person struct {
	Name string
	Age  int
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


}
