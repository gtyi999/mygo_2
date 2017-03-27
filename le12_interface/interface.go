package main

import (
	"fmt"
)

type TZ int
func (tz* TZ)Increase(num int){
	*tz+= TZ(100)
}

type  Connecter interface {
	Connect()
}
type  USB interface {
	Name() string
	Connecter
}

type PhoneConnecter struct {
	name string
}
func (pc PhoneConnecter)Name() string{
	return pc.name
}
func (pc PhoneConnecter)Connect(){
	fmt.Println("Connect:",pc.name)
}

func DisConnected(usb interface{}){
	/*if pc,ok:=usb.(PhoneConnecter);ok{
		fmt.Println("Disconnected:",pc.name)
		return
	}
	fmt.Println("Unknown decive")*/

	switch v:=usb.(type){
	case PhoneConnecter:
		fmt.Println("Disconnected:",v.name)
	default:
		fmt.Println("Unknown device.")
	}
}

type TVConnecter struct {
	name string
}
func (tv TVConnecter)Connect(){
	fmt.Println("Connected:",tv.name)
}

func main(){
	fmt.Println("1-课堂作业 start")
	var a TZ
	a.Increase(100)
	fmt.Println(a)
	fmt.Println("1-课堂作业 end")
	fmt.Println("2--接口的定义与声明 start")
	var b USB
	b=PhoneConnecter{"PhoneConecter"}
	b.Connect()
	DisConnected(b)
	fmt.Println("2--接口的定义与声明 end")
	fmt.Println("3--嵌入接口 start")
	fmt.Println("3--嵌入接口 end")

	fmt.Println("4--接口的类型转换 start")
	pc2:=PhoneConnecter{"PhoneConnecter"}
	var i1 Connecter
	i1=Connecter(pc2)
	i1.Connect()
	fmt.Println("4--接口的类型转换 end")
	fmt.Println("5---tvConnecter start")
	//tv:=TVConnecter{"TVConnecter"}
	//var i2 USB
	//i2=USB(tv)//接口只能从大向小转换，不能从小向大转换cannot convert tv (type TVConnecter) to type USB
	fmt.Println("5---tvConnecter end")

	fmt.Println("6--复制品不能修改原来的 start")
	pc6:=PhoneConnecter{"PhoneConnecter"}
	var a6 Connecter
	a6=Connecter(pc6)
	a6.Connect()
	pc6.name="ok"
	a6.Connect()
	fmt.Println("6---复制品不能修改原来的 end")
	fmt.Println("7--只有当接口存储类型和对象都为nil时,接口才等于nil start")
	var a7 interface{}
	fmt.Println(a7==nil)
	var p7 *int=nil
	a7=p7
	fmt.Println(a7==nil)

	//8--接口调用不会做receiver的自动转换
	//9--接口同样支持匿名字段方法
	//10--接口也可以实现类似OOP的多态
	//11--空接口可以作为任何类型数据的容器

}
