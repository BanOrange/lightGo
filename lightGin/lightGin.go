package lightGin

import (
	"net/http"
	"log"
)

//开始进行分组路由的构建
type RouterGroup struct {
	prefix string
	middlewares []HandlerFunc //对中间件的支持
	parent *RouterGroup
	engine *Engine
}
//请求处理函数的定义
type HandlerFunc func(*Context)

//通过阅读net中处理器函数的方法，我们需要实现ServeHTTP的接口
//并且再此基础上实现动态路由，那么就需要记录路由表
type Engine struct {
	router *router
	*RouterGroup
	groups []*RouterGroup //存储所有的路由分组
}

//构造函数,返回一个Engine实例指针
func New() *Engine {

	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

//为group再添加新的group，并且存储到engine中
func (group *RouterGroup) Group(prefix string) *RouterGroup{
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups,newGroup)
	return newGroup
}

//通过routerGroup来添加路由，此时会先将前缀加到pattern中，再加入到router中
func (group *RouterGroup) addRoute(method string,comp string,handler HandlerFunc){
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s",method,pattern)
	group.engine.router.addRoute(method,pattern,handler)
}

func (group *RouterGroup) GET(pattern string,handler HandlerFunc){
	group.addRoute("GET",pattern,handler)
}

func (group *RouterGroup) POST(pattern string,handler HandlerFunc){
	group.addRoute("POST",pattern,handler)
}

func(engine *Engine) addRoute(method string,pattern string,handler HandlerFunc){
	engine.router.addRoute(method,pattern,handler)

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
	c := newContext(w,req)
	engine.router.handle(c)
}


