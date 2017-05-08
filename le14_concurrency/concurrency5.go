package main

import (
	//"fmt"

)
import (
	"fmt"
	"time"
)

func main(){
	//3--可用空的select来阻塞main函数
	//4--可设置超时
	c:=make(chan  bool)

        select {
	case v:=<-c:
	fmt.Println(v)
	case <-time.After(3*time.Second):
	fmt.Println("timeout")
	}

}

