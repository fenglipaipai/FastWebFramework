package middleware

import (
	"fmt"
	"github.com/hanzhongzi/FastWebFramework-new/framework"
)

func Test1() framework.ControllerHandler {
	//使用函数回调
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test1")
		c.Next() //调用Next往下调整，会自增context.index
		fmt.Println("middleware post test1")
		return nil
	}
}

func Test2() framework.ControllerHandler {
	//使用函数回调
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test2")
		c.Next() //调用Next往下调整，会自增context.index
		fmt.Println("middleware post test2")
		return nil
	}
}

func Test3() framework.ControllerHandler {
	//使用函数回调
	return func(c *framework.Context) error {
		fmt.Println("middleware pre test3")
		c.Next() //调用Next往下调整，会自增context.index
		fmt.Println("middleware post test3")
		return nil
	}
}
