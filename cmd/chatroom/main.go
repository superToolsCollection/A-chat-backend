package main

import (
	"A-chat-backend/pkg/file"
	"A-chat-backend/server"
	"fmt"
	"log"
	"net/http"
)

/**
* @Author: super
* @Date: 2020-09-23 07:29
* @Description:
**/

var (
	addr = ":2022"
)

//func init() {
//	global.Init()
//}

func main() {
	s, err := file.ReadFile("/Users/super/develop/A-chat-backend/cmd/chatroom/banner.txt")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(s)

	server.RegisterHandle()

	log.Fatal(http.ListenAndServe(addr, nil))
}
