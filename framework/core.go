package framework

import "net/http"

// 框架核心结构
// type Core struct{

// }

// func NewCore() *Core {
// 	return &Core{}
// }


// func (c *Core) ServerHTTP(response http.ResponseWriter, request *http.Request){
// 	// TODO
// }


// ---------------
// 原版就是叫 servehttp 就没问题
type Core struct {
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO
}