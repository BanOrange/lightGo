package lightGin
//將router模塊拆解出來

import (
	"log"
	"net/http"
	"strings"
)
type router struct{
	handlers map[string]HandlerFunc
	roots map[string] *node
}

func newRouter() *router{
	return &router{
		roots: make(map[string] *node),
		handlers:make(map[string]HandlerFunc),
	}
}

//对路由地址进行拆分parts，只有*是特殊允许的符号
func parsePattern(pattern string) []string{
	vs := strings.Split(pattern,"/")

	parts := make([]string,0)
	for _,item := range vs{
		if item != "" {
			parts = append(parts,item)
			if item[0] == '*'{
				break;
			}
		}
	}
	return parts
}
func (r *router) addRoute(method string,pattern string,handler HandlerFunc){
	parts := parsePattern(pattern)

	log.Printf("Route %4s-%s",method,pattern)
	key := method + "-" + pattern

	_,ok := r.roots[method]
	if !ok{
		r.roots[method] = &node{}
	}

	r.roots[method].insert(pattern,parts,0)

	r.handlers[key] = handler
}

func (r *router) getRoute(method string,path string) (*node,map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root,ok := r.roots[method]
	//当前方法的前缀树还没有创建出来
	if !ok{
		return nil,nil
	}

	n:=root.search(searchParts,0)
	//能够查询到对应结果，那么就需要将其中的:和*进行处理
	if n!=nil{
		parts := parsePattern(n.pattern)
		
		for index,part := range parts{
			if part[0] == ':'{
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) >1{
				params[part[1:]] = strings.Join(searchParts[index:],"/")
				break;
			}

		}
		return n,params
	}
	return nil,nil
}

func (r *router) handle(c *Context){
	n,params := r.getRoute(c.Method,c.Path)
	if n!=nil{
		//这里主要是为了让handler能从context中获取到路径参数信息
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers,r.handlers[key])
	}else{
		c.handlers = append(c.handlers,func(c *Context){
			c.String(http.StatusNotFound,"404 NOT FOUND:%s\n",c.Path)
		})
	}
	c.Next()
}