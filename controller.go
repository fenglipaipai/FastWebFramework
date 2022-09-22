package main

import (
	"context"
	"fmt"
	"github.com/fenglipaipai/FastWebFramework/framework"
	"log"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	// 这个 channal 负责通知结束
	finish := make(chan struct{}, 1)
	// 这个 channel 负责通知 panic 异常
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(5*time.Second))
	// 这里记得当所有事情处理结束后调用 cancel，告知 durationCtx 的后续 Context 结束
	defer cancel()

	// mu := sync.Mutex{}
	go func() {
		// 这里增加异常处理
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		//DO real action
		// 这里做具体的业务
		time.Sleep(10 * time.Second)
		c.Json(200, "ok")
		finish <- struct{}{}

	}()
	//请求监听的时候增加锁机制
	select {
	case p := <-panicChan: // 监听 panic 异常事件
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Panicln(p)
		c.Json(500, "panic")

	case <-finish: // 新的 goroutine 结束的时候通过一个 finish 通道告知父 goroutine // 监听结束事件
		fmt.Println("finish")
	case <-durationCtx.Done(): // 监听超时事件
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.Json(500, "time out")
		c.SetHasTimeout()
	}
	return nil
}
