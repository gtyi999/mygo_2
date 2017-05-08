package main

import (
	"fmt"
	"net/http"
	"io"
	"log"
)

func main(){
	//设置路由
	fmt.Println("in to main")
	http.HandleFunc("/",SayHello)
	err:=http.ListenAndServe("127.0.0.1:9084",nil)
	if err!=nil{
		log.Fatal(err)
	}
}

func SayHello(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Hello world,this is version 1")
}
