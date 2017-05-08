package main

import (
	"fmt"
)


func main(){
	fmt.Println("hello main")

	var entityIDs[] string
	entityIDs=append(entityIDs,"1")
	entityIDs=append(entityIDs,"1")
	entityIDs=append(entityIDs,"2")
	entityIDs=append(entityIDs,"3")
	entityIDs=append(entityIDs,"3")

	var pc map[string] int
	pc = make(map[string] int)
	for _,v:=range entityIDs {
		if _, ok := pc[v]; ok {
			pc[v]=pc[v]+1
		}else{
			pc[v]=1
		}
	}

	for k,v:=range pc{
		fmt.Println(k,v)
	}
	//var pc map[string] int
	//pc = make(map[string] int)
	//pc["1"]= 0
	//pc["1"] = 0
	//pc["2"] = 0

	//for k,v:=range pc{
	//	fmt.Println(k,v)

		//if _, ok := map[k]; ok {
		//	//存在}
		//}
		//delete(pc,"qingdao")
		//qingdao,ok := pc["qingdao"]
		//if ok{
		//	fmt.Println(qingdao)
		//}else{
		//	fmt.Println("元素不存在")
		//}
	//}
}
