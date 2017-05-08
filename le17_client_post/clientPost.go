package main

import (
	"fmt"
	"bytes"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"time"
)


type PushInput struct {
	Type      		int  `json:"type" form:"type" `
	MessagID      	string  `json:"message_id" form:"message_id" `
	BatchID      	int  `json:"batch_id" form:"batch_id" `
	AopsID      	[]string  `json:"aops_id" form:"aops_id" `
	MessagCreateTime      	int  `json:"message_create_time" form:"message_create_time" binding:"required"`
	MessagDeadLine      	int  `json:"message_deadline" form:"message_deadline" `
	MessageParams               string `json:"message_params" form:"message_params" `
	MessageTemplateID        int  `json:"message_template_id" form:"message_template_id" binding:"required"`
	MessageUrlParams          string  `json:"message_url_params" form:"message_url_params"`
	YMRedurl               string `json:"ymredirectionurl" form:"ymredirectionurl"`
	IsNotify	   int     `json:"isnotify" form:"isnotify"`
	Iddatas           []IDdata `json:"id_data" form:"id_data"`
}
type IDdata  struct {
	AopsID      	string  `json:"aops_id" form:"aops_id" `
	EntityID      	string  `json:"entity_id" form:"entity_id" `
}

func main(){
	fmt.Println("hello main")
        //{"message_params":"11","type":0,"message_create_time":1493281310,"message_url_params":"aaaaaaaaaaaa","aops_id":["4116618"],"message_template_id":3402}
	var in PushInput
	var aopsid []string
	aopsid=append(aopsid,"4116618")

	in.AopsID=aopsid
	in.MessageParams="11"
	in.MessageUrlParams="aaaaaaaaaaaa"
	in.Type=0
	in.MessageTemplateID=3403
	in.MessagCreateTime=int(time.Now().Unix())


	b, err := json.Marshal(in)
	if err != nil {
		fmt.Println("json err:", err)
	}

	body := bytes.NewBuffer([]byte(b))
	res,err := http.Post("http://119.18.231.128:8010/msg/push", "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", result)
}
