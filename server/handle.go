package server

import (
	"A-chat-backend/logic"
	"net/http"
)

/**
* @Author: super
* @Date: 2020-09-23 07:30
* @Description:
**/

func RegisterHandle() {
	// 广播消息处理
	go logic.Broadcaster.Start()

	http.HandleFunc("/", homeHandleFunc)
	http.HandleFunc("/user_list", userListHandleFunc)
	http.HandleFunc("/ws", WebSocketHandleFunc)
}