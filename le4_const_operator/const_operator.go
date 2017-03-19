package main

import (
	"fmt"
)
//1---常量的声明与定义
//const a int =1
//const b ='A'
//2----常量的声明与定义
/*const(
	c=1
	d=2
	e=3
)*/

//3--常量的声明与定义
/*const(
	c=a+1
	d=a+2
	e=a+3
)*/
//4---常量的声明与定义
//const zz1,zz2,zz3=1,"2",'c'

//5---常量的声明与定义
//const(
//	aa1=1
//	aa2
//	aa3
//)
//6---常量的声明
const (
	bb1,bb2=1,"2"
	bb3,bb4//上面如果成对的声明与定义，下面这里也是成对的，并且初值也是一样的。
)
//7---枚举
const(
	cc1='A'
	cc2=iota
	cc3='B'
	cc4=iota
)
//8--运算符
/*
6: 0110
11:1011
------------------
&  0010
|  1111
^  1101
&^ 0100//关键看11的二进制位，如果为第一位为1，则将结果的第一位置为0,如果第一位为0，刚将6的第一位取为结果
 */

//9--实现计算的kb,MB,G的单位的枚举
const (
	B float64 =1<<(iota *10)
	KB
	MB
	GB

)

func main(){
        fmt.Println("start")
	//fmt.Println(zz1)//输出1
	//fmt.Println(zz2)//输出2
	//fmt.Println(zz3)//输出99

	//fmt.Println(aa1)
	//fmt.Println(aa2)//输出1
	//fmt.Println(aa3)//输出1，如是变量组的形式赋值，则以后面的变量都默认为第一个变量的值

	//fmt.Println(bb1)
	//fmt.Println(bb2)
	//fmt.Println(bb3)//上面如果成对的声明与定义，下面这里也是成对的，并且初值也是一样的。
	//fmt.Println(bb4)//上面如果成对的声明与定义，下面这里也是成对的，并且初值也是一样的。

	//fmt.Println(cc1)
	//fmt.Println(cc2)//iota输出1
	//fmt.Println(cc3)//输出2
	//fmt.Println(cc4)//输出3

	//fmt.Println(1^2)//0001 ^ 0010
        //fmt.Println(!true)
	//fmt.Println(1<<10<<10>>10)//1024

	//fmt.Println(6 & 11)//2
	//fmt.Println(6 | 11)//15
	//fmt.Println(6 ^ 11)//13
	//fmt.Println(6 &^ 11)//4

	dd1:=1
	if dd1>0 && (10/dd1)>1{
		fmt.Println("ok")
	}else {
		fmt.Println("dd1 is zero")
		return
}
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

}
