package main

import (
	"fmt"
)

func main() {
	fmt.Println("array start")
	var array1 [2]int
	fmt.Println(array1[0])//[0 0]
	//1--数组非法操作1
	//var arr2 [1]int
	//array1=arr2//数组的长度不同不能作赋值操作
	var arr2 [2]int
	array1 = arr2//长度相同，可以赋值
	fmt.Println(array1)

	//2----string array
	fmt.Println("string1 array start")
	var arrstr1 [2]string
	fmt.Println(arrstr1)
	fmt.Println("string1 end")

	//3---数组的初始化
	fmt.Println("arr3 --start")
	a := [2]int{1, 2}
	fmt.Println(a)
	fmt.Println("arr3--end")
	//4---数组初始化
	fmt.Println("array4---start")
	b := [20]int{19:1}
	fmt.Println(b)
	fmt.Println("array4--end")
	//5---数组的初始化
	fmt.Println("arr5--start")
	cc1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(cc1)
	fmt.Println("arr5--end")
	//6--array init
	fmt.Println("arr6--start")
	dd2 := [...]int{0:1, 1:2, 2:3}
	fmt.Println(dd2)
	fmt.Println("arr6--end")
	//7---指向数组的指针
	fmt.Println("arry7_pointer start")
	ee2 := [...]int{99:1}
	var p *[100]int = &ee2
	fmt.Println(p)
	fmt.Println("array8_pointer end")
	//8---指针数组
	fmt.Println("pointer_array start")
	x, y := 1, 2
	ff2 := [...]*int{&x, &y}
	fmt.Println(ff2)
	fmt.Println("pointer_aray_end")
	//9---数组可以用==和!=比较
	fmt.Println("9--arr_compare")
	aaa := [2]int{1, 2}
	bbb := [2]int{1, 2}
	fmt.Println(aaa == bbb)
	ccc := [2]int{1, 3}
	fmt.Println(aaa == ccc)
	//ddd:=[1]int {1}
	//fmt.Println(aaa==ddd)//这个不同长度的数组不能比较
	fmt.Println("9--arr_compare end")
	//10--指向数组的指针
	fmt.Println("10pointer_array start")
	p1 := new([10]int)
	p1[1] = 2
	fmt.Println(p1)
	fmt.Println("-------------")
	aaa2 := [10]int{}
	aaa2[1] = 2
	fmt.Println(aaa2)
	fmt.Println("10pointer_array end")
	//11--多维数组
	fmt.Println("11Multi_array")
	q1 := [2][3]int{{1, 1, 1}, {2, 2, 2}}
	fmt.Println(q1)
	q2 := [2][3]int{{1:1}, {2:2}}
	fmt.Println(q2)
	fmt.Println("11Multi_array end")
	//12---go的冒泡排序
	fmt.Println("12Bubble sort start")
	q3 := [...]int{5, 2, 6, 3, 9}
	fmt.Println(q3)
	num := len(q3)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if q3[i] < q3[j] {
				temp := q3[i]
				q3[i] = q3[j]
				q3[j] = temp
			}
		}
	}
	fmt.Println(q3)
}