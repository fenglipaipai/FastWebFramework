package main

import (
	"fmt"
	"net/http"
)
import "github.com/fenglipaipai/FastWebFramework/framework"

func main() {
	core := framework.NewCore()
	registerRouters(core)
	server := &http.Server{
		//请求核心处理函数
		Handler: core,
		Addr:    ":8080",
	}
	fmt.Println("server is running.....")

	server.ListenAndServe()
}
