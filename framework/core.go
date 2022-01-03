package framework

import (
	"log"
	"net/http"
)

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
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{
		router: map[string]ControllerHandler{},
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

// 这个是 handler 调用的 serveHTTP
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// 调试打印
	// fmt.Println("servehttp")

	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	router := c.router["foo"]
	if router == nil {
		return
	}

	log.Println("core.router")

	router(ctx)
}
