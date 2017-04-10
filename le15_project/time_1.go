package main

import (
	"fmt"
	"time"
)


func main(){
	//1--Mon Jan _2 15:04:05 2006只能使用这个固定的常量,否则取出的时间不准确
	t:=time.Now()
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t.Format("2006-01-02 15：04：05"))




	//3 格式串类型
	//当然如果上面没有你要的格式化类型，那就看下面的附表：
	//const (
	//	ANSIC       = "Mon Jan _2 15:04:05 2006"
	//	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	//	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	//	RFC822      = "02 Jan 06 15:04 MST"
	//
	//	// RFC822 with numeric zone RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	//	RFC822Z     = "02 Jan 06 15:04 -0700"
	//	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	//
	//	// RFC1123 with numeric zone RFC3339     = "2006-01-02T15:04:05Z07:00"
	//	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
	//	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	//
	//	// Handy time stamps. Stamp      = "Jan _2 15:04:05"
	//	Kitchen     = "3:04PM"
	//	StampMilli = "Jan _2 15:04:05.000"
	//	StampMicro = "Jan _2 15:04:05.000000"
	//	StampNano  = "Jan _2 15:04:05.000000000"
	//)
}



