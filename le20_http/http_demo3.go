package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter,*http.Request)
func main(){

	mux=make(map[string]func(http.ResponseWriter,* http.Request))
	mux["/hello"]=SayHello
	mux["/bye"]=SayBye

	server:=&http.Server{
		Addr:"127.0.0.1:9084",
		Handler: &myHandler{},
		ReadTimeout: 5*time.Second,
	}

	err:=server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}

type myHandler struct {

}

func (* myHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	//io.WriteString(w,"URL: "+r.URL.String())
	if h,ok:=mux[r.URL.String()]; ok {
		h(w,r)
		return
	}
}



func SayHello(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Hello world,this is version 3")
}

func SayBye(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"bye bye ,version 3")
}
