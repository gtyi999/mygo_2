package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age  int
}

func (u User) Hello(){
	fmt.Println("Hello world!")
}

type Mannager struct {
	User
	title string
}

func (u User)Hello5(name string){
	fmt.Println("Hello",name,", my name is ",u.Name)
}
func main(){
	u:=User{1,"ok",12}
	Info(u)


	fmt.Println("2---reflection start")
	m:=Mannager{User:User{1,"ok",12},title:"123"}
	t2:=reflect.TypeOf(m)
	fmt.Printf("%#v\n",t2.Field(1))
	fmt.Printf("%#v\n",t2.FieldByIndex([]int{0,1}))
	fmt.Println("2---reflection end")

	fmt.Println("3--设置值 start")
	x3:=123
	v3:=reflect.ValueOf(&x3)
	v3.Elem().SetInt(999)
	fmt.Println(x3)
	fmt.Println("3--设置值 end")

	fmt.Println("4--修改结构的值 start")
	u4:=User{1,"ok",12}
	Set(&u4)
	fmt.Println(u4)
	fmt.Println("4--修改结构的值 end")

	fmt.Println("5---start")
	u5:=User{1,"ok",12}
	u5.Hello5("joe")
	fmt.Println("5---end")
	fmt.Println("6---start")
	u6:=User{1,"ok",12}
	v6:=reflect.ValueOf(u6)
	mv6:=v6.MethodByName("Hello5")
	args:=[]reflect.Value{reflect.ValueOf("joe")}
	mv6.Call(args)
	fmt.Println("6---end")

}

func Info(o interface{}){
	fmt.Println("1---reflection start")
	t:=reflect.TypeOf(o)
	fmt.Println("Type:",t.Name())

	if k:=t.Kind();k!=reflect.Struct{
		fmt.Printf("XX")
		return
	}
	v:=reflect.ValueOf(o)
	fmt.Println("Fields:")
	for i:=0;i<t.NumField();i++{
		f:=t.Field(i)
		val:=v.Field(i).Interface()
		fmt.Printf("%6s:%v=%v\n",f.Name,f.Type,val)
	}

	for i:=0;i<t.NumMethod();i++{
		m:=t.Method(i)
		fmt.Printf("%6s:%v\n",m.Name,m.Type)
	}




}
func Set(o interface{}){
	v:=reflect.ValueOf(o)
	if v.Kind()==reflect.Ptr && !v.Elem().CanSet(){
		fmt.Println("XXXX")
		return
	}else{
		v=v.Elem()
	}

	f:=v.FieldByName("Name")
	if !f.IsValid(){
		fmt.Println("BAD")
		return
	}
	if  f.Kind()==reflect.String{
		f.SetString("BYEBYE")
	}
}

