package main

import(
	"fmt"
	"runtime"
	"sync"//同步包
)

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//c:=make(chan bool,10)
	wg:=sync.WaitGroup{}
	wg.Add(10)
	for i:=0;i<10;i++{
		//go Go(c,i)
		go Go(&wg,i)
	}
	wg.Wait()

}

func Go(wg *sync.WaitGroup,index int){
	a:=1
	for i:=0;i<1000000;i++{
		a+=i
	}
	fmt.Println(index,a)

	//c <- true
	wg.Done()

}
