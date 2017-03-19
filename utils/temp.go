package main

import(
	"fmt"
	//"math"
	"strconv"
)

type (
	byte int8
	rune int32
	文本 string
	bytesize int64
)
func main() {
	//1----数据类型的默认初始值
	//var a int
	//fmt.Println(a)//定义一个整型，初值输出为0

	//var a int32
	//fmt.Println(a)//定义一个int32整型，初值输出为0

	//var a float32
	//fmt.Println(a)//定义一个float32整型，初值输出为0

	//var a bool
	//fmt.Println(a)//定义一个bool整型，初值输出为false

	//var a string
	//fmt.Println(a)//定义一个string，初值输出为空

	//var a []int
	//fmt.Println(a)//定义一个int型数组，初值输出为空数组

	//var a [1]int
	//fmt.Println(a)//定义一个int型数组，初值输出为[0]

	//var a []bool
	//fmt.Println(a)//定义一个型数组，初值输出为空数组

	//var a [2]bool
	//fmt.Println(a)//定义一个int型数组，初值输出为[false false]

	//var a [1]byte
	//fmt.Println(a)//定义一个byte型数组，初值输出为[0]

	//2----math这个包记录的是数据类型的最小值和最大值
	//fmt.Println(math.MinInt8)//-128
	//fmt.Println(math.MaxInt8)//127
	//fmt.Println(math.MaxInt32)//2147483647

	//3--类型的别名
/*	var b 文本
	b="中文类型别名"
	fmt.Println(b)*/

	//4--变量的声明与赋值
	//var b int//第一种方法声明
	//b=1
	//c:=1//第二种隐式的转换声明c就是一个int
	//var d=123
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(d)
	//5---多个变量的声明与赋值
	//var a,b,c,d int=1,2,3,4//标准声明与赋值
	//var a,b,c,d =1,2,3,4//隐式的类型转换声明
	//a,b,c,d :=1,2,3,4//省略var关键字，隐式的类型转换声明
	//a,_,c,d :=1,2,3,4//空白的符号忽略,第二个参数不输出
	//fmt.Println(a,c,d)
	//6---变量的类型转换
	//var a float32 =1.1
	//fmt.Println(a)
	//b:=int(a)
	//fmt.Println(b)

	//var c bool=true
	//d:=int(c)//Error:cannot convert c (type bool) to type int
	//fmt.Println(d)
	//var a int
	//a=1
	//fmt.Println(a)

	//7--课堂作业
	//var a int =65
	//b:=string(a)
	//fmt.Println(b)//输出A,就是ACCII码

	//var a int=65
	//b:=strconv.Itoa(a)//相当于整形转字符串
	//fmt.Println(b)//输出65字符串
}

