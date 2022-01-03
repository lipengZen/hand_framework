package main

import (
	"hand_by_hand_framework/framework"
	"net/http"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		// 自定义的请求核心处理函数
		Handler: core, //framework.NewCore(),
		// 请求监听地址
		Addr: "localhost:8080",
	}
	server.ListenAndServe()
	// http.FileServer()

}
