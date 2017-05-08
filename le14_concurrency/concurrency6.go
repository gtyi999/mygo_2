package main

import (
	"fmt"
	"time"
)

func ready(w string,sec int64){
	time.Sleep(time.Duration(sec * 1e9))
	fmt.Println(w,"is ready")
}

func main(){
	go ready("Tee",2)
	go ready("Coffee",1)

	fmt.Println("I m waiting")
	time.Sleep(5*1e9)
}


