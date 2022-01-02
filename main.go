package main

import (
	"hand_by_hand_framework/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: framework.NewCore(),
		// 请求监听地址
		Addr: ":8080",
	}
	server.ListenAndServe()
	// http.FileServer()
}
