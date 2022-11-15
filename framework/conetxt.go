package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

//自定义 conetxt

type Context struct {
	request        *http.Request       //封装以后实现
	responseWriter http.ResponseWriter //封装以后实现
	ctx            context.Context
	hasTimeout     bool                //是否超时标记
	writeMux       *sync.Mutex         //写保护机制
	handlers       []ControllerHandler //当前请求的handler链条
	index          int                 //当前请求调用到调用链的哪个节点
	params         map[string]string   // url路由匹配的参数
}

//控制调用context的下一个函数
func (ctx *Context) Next() error {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx); err != nil {
			return err
		}
	}
	return nil
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writeMux:       &sync.Mutex{},
		handlers:       []ControllerHandler{},
		index:          -1, //这样才能保证第一次调用的时候index为0
	}
}

// #region base function
//base 封装基本函数功能，比如获取http.Request结构

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writeMux
}
func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}
func (ctx *Context) GetResponse() http.ResponseWriter {
	return ctx.responseWriter
}
func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}
func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

// #endregion

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

// #region context function   implement context.Context
// context 实现标准Context接口

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}
func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}
func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}
func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.BaseContext().Value(key)
}

// #endregion

// #region request function  query url
//request 封装了http.Request 的对外接口
/*func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			intval, err := strconv.Atoi(vals[len-1])
			if err != nil {
				return def
			}
			return intval
		}
	}
	return def
}
func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return vals[len-1]
		}
	}
	return def
}
*/
func (ctx *Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

/*
func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.URL.Query())
	}
	return map[string][]string{}
} */

// #endregion

//  #region request function   post
// request 封装了http.Request 的对外接口
func (ctx *Context) FromAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.PostForm)
	}
	return map[string][]string{}
}

/*
func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FromAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			intval, err := strconv.Atoi(vals[len-1])
			if err != nil {
				return def
			}
			return intval
		}
	}
	return def

}
func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FromAll()
	if vals, ok := params[key]; ok {
		len := len(vals)
		if len > 0 {
			return vals[len-1]
		}
	}
	return def
} */
func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FromAll()
	if vals, ok := params[key]; ok {
		return vals
	}
	return def
}

// #endregion

// #region application/json post
/*
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
} */

// #endregion

// #region response
/*
func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.Write(byt)
	return nil

}
*/
func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

/*
func (ctx *Context) Text(status int, obj string) error {
	return nil
}*/

// #endregion

// 为context设置handlers
func (ctx *Context) SetHandlers(handlers []ControllerHandler) {
	ctx.handlers = handlers
}

// 设置参数
func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}
