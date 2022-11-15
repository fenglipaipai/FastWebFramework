package main

import (
	"github.com/hanzhongzi/FastWebFramework-new/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("/user/login", "def")
	// 等待10s才结束执行
	time.Sleep(1 * time.Second)
	// 输出结果
	c.SetOkStatus().Json("ok, UserLoginController: " + foo)
	return nil
}

func SubjectNameController(c *framework.Context) error {
	//打印控制器名字

	c.SetOkStatus().Json("ok, SubjectNameController")
	return nil
}
func SubjectDelController(c *framework.Context) error {
	//打印控制器名字
	c.SetOkStatus().Json("ok, SubjectDelController")
	return nil
}
func SubjectUpdateController(c *framework.Context) error {
	//打印控制器名字
	c.SetOkStatus().Json("ok, SubjectUpdateController")
	return nil
}
func SubjectGetController(c *framework.Context) error {
	//打印控制器名字
	c.SetOkStatus().Json("ok, SubjectGetController")
	return nil
}
func SubjectListController(c *framework.Context) error {
	//打印控制器名字
	c.SetOkStatus().Json("ok, SubjectListController")
	return nil
}
