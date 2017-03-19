package main

import (
	"fmt"
	"sort"
)

func main(){
	//1--定义map
	fmt.Println("map1----")
	var m map[int]string
	m=make(map[int]string)
	fmt.Println(m)
	//2---第二种定义方式
	var m2 map[int]string=make(map[int]string)
	fmt.Println(m2)
	//3---map定义方式3
	m3:=make(map[int]string)
	fmt.Println(m3)
	//4--map的初始化
	m[1]="ok"
	fmt.Println(m)
	//5，取出键的值
	a:=m[1]
	fmt.Println(a)
	//6,删除键值
	fmt.Println("delete----")
	delete(m,1)
	b:=m[1]
	fmt.Println(b)//输出为空,成功删除
	fmt.Println("delete--end")
	//7--复杂的map结构
	var m4 map[int]map[int]string
	m4=make(map[int]map[int]string)
	a,ok:=m4[2][1]
	if !ok{
		m4[2]=make(map[int]string)
	}
	m4[2][1]="GOOD"
	a=m4[2][1]
	a,ok=m4[2][1]//再取一次，变为true
	fmt.Println(a,ok)
	//8--迭代操作
	slice_1:=make([]map[int]string,5)
	for _,v:=range slice_1{
		v=make(map[int]string)
		v[1]="ok"
		fmt.Println(v)

	}
	fmt.Println(slice_1)
	//9--迭代操作2
	slice_2:=make([]map[int]string,5)
	for i:=range slice_2{
		slice_2[i]=make(map[int]string)
		slice_2[i][1]="ok"
		fmt.Println(slice_2[i])

	}
	fmt.Println(slice_2)
	//10--map的key间接排序
	m1:=map[int]string{1:"a",2:"b",3:"c",4:"d",5:"e"}
	slic_3:=make([]int,len(m1))
	i:=0
	for k,_:=range m1{
		slic_3[i]=k
		i++
	}
	sort.Ints(slic_3)
	fmt.Println(slic_3)
	//11--课堂作业,键值对调
	m5:=map[int]string{1:"a",2:"b",3:"c"}
	m6:=make(map[string]int)
	for k,v:=range m5{
		m6[v]=k
	}
	fmt.Println(m6)


}


