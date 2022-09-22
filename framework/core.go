package framework

import (
	"log"
	"net/http"
	"strings"
)

// Core  框架核心结构
type Core struct {
	router map[string]*Tree // all routers
}

// NewCore 初始化框架核心结构
func NewCore() *Core {
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	router["PUT"] = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}
func (c *Core) Group(prefix string) IGroup {
	return NewGr192.168.9.15192.168.9.15480005
	480005oup(c, prefix)
}

// http method wrap

//匹配GET方法，增加路由规则
func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatalf("add router error: ", err)
	}
}

//匹配 POST 方法，增加路由规则
func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatalf("add router error: ", err)
	}
}

//匹配PUT方法，增加路由规则
func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatalf("add router error: ", err)
	}
}

// 匹配DELETE方法，增加路由规则
func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatalf("add router error: ", err)
	}
}

// http method wrap end

/*
// 匹配路由，如果没有匹配到 返回nil
func (c *Core) FindRoutByRequest(request *http.Request) ControllerHandler {
	//uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	//查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
} */

// 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	// uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}

// ServerHTTP 框架核心结构实现Handler接口
// 所有的请求都进入这个函数，这个函数负责路由分发
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	//log.Panicln("core.serverHTTP")
	//封装自定义context
	ctx := NewContext(request, response)
	// 一个简单的路由选择器，这里直接写死为测试路由foo
	router := c.FindRouteByRequest(request)

	if router == nil {
		//没找到就打印日志
		ctx.Json(404, "not found")
		return
	}

	//调用路由函数，如果然后err代表内部存在错误，返回500
	if err := router(ctx); err != nil {
		ctx.Json(500, "inner error")
		return
	}
}
