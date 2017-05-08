package main

import (
	//    "bytes"
	//    "encoding/json"
	"fmt"
	//    "os"
	//    "gopkg.in/mgo.v2"
	//    "gopkg.in/mgo.v2/bson"
	//    "io/ioutil"
	//    "net/http"
	//    "strings"
	//    "net/url"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	for i := 0; i < 10; i++ {
		time := <-ticker.C
		fmt.Println(time.String())
	}

	fmt.Println("hhh")
}