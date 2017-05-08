package main

import (
	"fmt"
	"net/http"
	"io"
	"log"
	"os"
)

func main(){
	fmt.Println("hello world")
	mux:=http.NewServeMux()
	mux.Handle("/",&myHandler{})

	mux.HandleFunc("/hello",SayHello)

	wd,err:=os.Getwd()
	if err!=nil{
		log.Fatal(err)
	}

	mux.Handle("/static",http.StripPrefix("/static/",http.FileServer(http.Dir())))




	err:=http.ListenAndServe("127.0.0.1:9084",mux)
	if err!=nil{
		log.Fatal(err)
	}
}

type myHandler struct {

}

func (*myHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"URL"+r.URL.String())
}
func SayHello(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Hello world,this is version 2")
}