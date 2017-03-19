package main

import (
	"fmt"
	//"strconv"
)

//
func main(){
	fmt.Println("start")
	//1---++运算符用法
	a:=1
	//a:=a++ //不能这样写，但C++这样是可以的
	a++//go只能这样递增
	//2---指针用法
	var p *int=&a
	fmt.Println(p)//输出p的地址0xc04203e1e8
	fmt.Println(*p)//1
	//3--if语句,b的作用域只在if语句块中
	b:=10
	if b:=1; b>0{
		fmt.Println(b)//b的作用域在if语句块中
	}else{
		fmt.Println("b<1")
	}
	fmt.Println(b)//此时b为外部的b
	//4--if语句
	if c,d:=1,2;c>0{
		fmt.Println(c)
		fmt.Println(d)
	}else{
		fmt.Println("c<0")
	}
	//5---循环语句for，在go里只有for，没有while语句，左大括号必须和for在同一行
	fmt.Println("for start")
	cc1:=0
	for{
		cc1++
		if cc1>3{
			break
		}
		fmt.Println(cc1)
	}
	//6---for
	fmt.Println("for2")
	dd1:=0
	for dd1<=3{
		dd1++
		if dd1>3{
			break
		}
		fmt.Println(dd1)
	}
	fmt.Println("over for2")
	//7---for3
	fmt.Println("for3")
	ee1:=0
	for i:=0;i<3;i++{
		ee1++
		fmt.Println(ee1)
	}
	fmt.Println("for3 over")
	//8---for4
	fmt.Println("for4 start")
	str:="abcde"
	ilen:=len(str)
	for i:=0;i<ilen ;i++  {
		fmt.Println(str)
	}
	fmt.Println("for4 end")
	//9---switch
	fmt.Println("switch1 start")
	ff1:=1;
	switch ff1 {
	case 0:
		fmt.Println("ff1==0")
	case 1:
		fmt.Println("ff1==1")
	case 2:
		fmt.Println("ff1==2")
	}
	fmt.Println("switch1 end")

	fmt.Println("switch2 start")
	gg:=1
	switch  {
	case gg>=0:
		fmt.Println("gg==0")
		fallthrough
	case gg>=1:
		fmt.Println("gg==1")
	default:
		fmt.Println("None")
	}
	fmt.Println("switch2 end")

	fmt.Println("switch3 start")
	switch hh:=1;{
	case hh>=0:
		fmt.Println("hh>=0")
		fallthrough
	case hh>=1:
		fmt.Println("hh>=1")
	default:
		fmt.Println("None")
	}
	fmt.Println("swwitch4 end")
	//10---跳转语句
	fmt.Println("break1 start")
	LABEL1:
	for  {
		for i:=0;i<10;i++{
			if i>3{
				break LABEL1
			}
		}

	}
	fmt.Println("break1 end")
        //11---goto语句
	fmt.Println("goto1 start")
	for {
		for i:=0;i<10;i++{
			if i>3 {
				goto LABLE2
			}
		}
	}
	LABLE2:
	fmt.Println("goto1 end")
	//12---continue
	fmt.Println("continue start")
	LABLE3:
	for i:=0;i<10;i++ {
		for {
			continue LABLE3
			fmt.Println(i)
		}
	}
	fmt.Println("continue end")

	//13---课程作业,死循环
	fmt.Println("start")
	LABLE4:
	for i:=0;i<10;i++ {
		for {
			fmt.Println(i)
			goto LABLE4
		}
	}
	fmt.Println("end")


}
