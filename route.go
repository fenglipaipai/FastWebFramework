package main

import (
	"github.com/fenglipaipai/FastWebFramework/framework"
	"time"
)

// 注册路由
func registerRouters(core *framework.Core) {
	/*
		需求 1：HTTP 方法匹配
		需求 2：静态路由匹配
		需求 3：批量通用前缀
		需求 4：动态路由匹配
	*/

	// core.Get("foo", framework.TimeoutHandler(FooControllerHandler, time.Second*1))
	core.Get("/foo", FooControllerHandler)
	// 需求1+2:HTTP方法+静态路由匹配
	core.Get("/user/login", framework.TimeoutHandler(UserLoginController, time.Second))

	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 需求4:动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}

}
