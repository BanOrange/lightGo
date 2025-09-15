package lightGin

import (
	"net/http"
	"fmt"
)

//请求处理函数的定义
type HandlerFunc func(http.ResponseWriter, *http.Request)

//通过阅读net中处理器函数的方法，我们需要实现ServeHTTP的接口
//并且再此基础上实现动态路由，那么就需要记录路由表
type Engine struct {
	router map[string]HandlerFunc
}

//构造函数,返回一个Engine实例指针
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}

}

func(engine *Engine) addRoute(method string,pattern string,handler HandlerFunc){
	key := method + "-" + pattern
	engine.router[key] = handler

}

//下面对不同的请求进行处理
func(engine *Engine) POST(pattern string,handler HandlerFunc) {
	engine.addRoute("POST",pattern,handler)
}

func(engine *Engine) GET(pattern string,handler HandlerFunc) {
	engine.addRoute("GET",pattern,handler)
}

//开启服务器
func (engine *Engine) Run(addr string) (err error){
	return http.ListenAndServe(addr,engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
	key := req.Method + "-" + req.URL.Path
	if handler,ok := engine.router[key]; ok {
		handler(w,req)
	}else{
		fmt.Fprintf(w,"404 NOT FOUND: %s\n",req.URL)
	}
}


