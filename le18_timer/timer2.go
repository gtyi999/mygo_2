package main

import (
	"time"
	"fmt"
)

func startTimer() {
	go func() {
		for {
			now := time.Now()

			//next := time.Date(now .Year(), now .Month(), now .Day(), now.Hour(), now.Minute()+1, 25, 0, now .Location())
			next := time.Date(now .Year(), now .Month(), now .Day(), 9, 30, 0, 0, now .Location())
			// 从程序开始到下一次需要执行的时间
			fmt.Println("下次执行时间还有:",next.Sub(now))
			t := time.NewTimer(next.Sub(now))
			<-t.C
			if next.Sub(now)>0  {
				fmt.Println("fisrt")
				//放业务函数
			}
			fmt.Println("当前时间:",time.Now().Hour(),":",time.Now().Minute(), " : ",time.Now().Second())
			//第二个运行时间点
			next = time.Date(now .Year(), now .Month(), now .Day(), 14, 30, 0, 0, now .Location())
			fmt.Println("下次执行时间还有:",next.Sub(now))
			t = time.NewTimer(next.Sub(now))
			<-t.C
			if next.Sub(now)>0 {
				fmt.Println("second")
				//放业务函数
			}
			fmt.Println("当前时间:",time.Now().Hour(),":",time.Now().Minute(), " : ",time.Now().Second())
			//第三个运行时间点
			next = time.Date(now .Year(), now .Month(), now .Day(), 19, 30, 0, 0, now .Location())
			fmt.Println("下次执行时间还有:",next.Sub(now))
			t = time.NewTimer(next.Sub(now))
			<-t.C
			if next.Sub(now)>0  {
				fmt.Println("three")
				//放业务函数
			}
			fmt.Println("当前时间:",time.Now().Hour(),":",time.Now().Minute(), " : ",time.Now().Second())

		}
	}()
	select {

	}
}
func main()  {
	startTimer()
}
