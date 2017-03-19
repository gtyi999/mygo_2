package main

import (
	"fmt"
)

func main(){
	fmt.Println("Slice1 create start")
	var a [10]int
	fmt.Println(a)

	s1:=a[9]
	fmt.Println(s1)
	fmt.Println("slice1 end")
	//2---slice的初始化
	fmt.Println("slice2 start")
	b:=[...]int {0,1,2,3,4,5,6,7,8,9}
	fmt.Println(b)
	s2:=b[5:10]
	fmt.Println(s2)
	fmt.Println("-------")
	s3:=b[5:len(b)]
	fmt.Println(s3)
	fmt.Println("---------")
	s4:=b[5:]
	fmt.Println(s4)
	fmt.Println("slice2 end")
	//3----slice的初始化
	fmt.Println("slice3--start")
	s5:=make([]int,3,10)//slice的长度为3，容量为10..如果超过容量刚再重新分配2倍的容量
	fmt.Println(s5)
	fmt.Println(len(s5),cap(s5))//3 10
	fmt.Println("-------------")
	s6:=make([]int,5)
	fmt.Println(len(s6),cap(s6))//5，5
	fmt.Println("slice3---end")
	fmt.Println("slice4---start--slice与底层数组的对应关系")
	a11:=[]byte {'a','b','c','d','e','f','g','h','i','j','k'}
	sa:=a11[2:5]
	fmt.Println(sa)
	fmt.Println(string(sa))
	fmt.Println("---------------")
	slice_b:=a11[3:5]
	fmt.Println(string(slice_b))
	fmt.Println("slice4---end--slice与底层数组的对应关系")
	fmt.Println("Reslice---start")
	sa2:=a11[2:5]//最大容量是9
	//sb2:=sa[9:11]//error panic: runtime error: slice bounds out of range
	sb2:=sa[1:3]
	fmt.Println(string(sa2))
	fmt.Println(len(sa2),cap(sa2))
	fmt.Println(string(sb2))
	fmt.Println("Reslice---end")
	//Append函数
	fmt.Println("Append---start")
	slice_11:=make([]int,3,6)
	fmt.Printf("%p\n",slice_11)
	slice_11=append(slice_11,1,2,3)
	fmt.Printf("%v %p\n",slice_11,slice_11)
	slice_11=append(slice_11,1,2,3)//超过了容量，从新分配
	fmt.Printf("%v %p\n",slice_11,slice_11)//可以看到地址的值和上面的不同了
	fmt.Println("-----------------")
	aaa1:=[]int {1,2,3,4,5}
	s_1:=aaa1[2:5]
	s_2:=aaa1[1:3]
	fmt.Println(aaa1)
	fmt.Println(s_1,s_2)
	//s_2=append(s_2,1,2,1,11,1,1,1,1,1,1,1)//当append重新分配时，s_1[0]=9的改变就不会影响s_2了
	s_1[0]=9//s_1和s_2相同的索引也会改变
	fmt.Println(s_1,s_2)
	fmt.Println("append---end")
	//copy函数
	fmt.Println("copy--start")
	s_3:=[]int {1,2,3,4,5,6}
	s_4:=[]int {7,8,9}
	copy(s_3,s_4)
	fmt.Println(s_3)//[7 8 9 4 5 6]
	fmt.Println("----------")
	s_5:=[]int {1,2,3,4,5,6}
	s_6:=[]int {7,8,9}
	copy(s_6,s_5)
	fmt.Println(s_6)//[1 2 3]
	fmt.Println("----------")
	s_7:=[]int {1,2,3,4,5,6}
	s_8:=[]int {7,8,9}
	copy(s_7[0:2],s_8[1:3])
	fmt.Println(s_7)//[8 9 3 4 5 6]
	fmt.Println("copy--end")
	//课堂作业
	s_9:=[]int {1,2,3,4}
	s_10:=s_9[:4]//拷贝全部s_9元素
	fmt.Println(s_10)
	s_11:=s_9[0:]
	fmt.Println(s_11)
	s_12:=s_9[:len(s_9)]
	fmt.Println(s_12)
	s_13:=s_9[:]
	fmt.Println(s_13)












}
