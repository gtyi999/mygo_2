package main

import (
	"fmt"
)

func main(){
	//1---都是输出c
	//s:=[]string{"a","b","c"}
	//for _,v:=range s{
	//	go func(){
	//		fmt.Println(v)
	//	}()
	//}
	//select {
	//}
	//输出如下
	//c
	//c
	//c

	//2--输出a,b,c
	s2:=[]string{"a","b","c"}
	for _,v:=range s2{
		go func(v string){
			fmt.Println(v)
		}(v)
	}

	select {
	}


}
