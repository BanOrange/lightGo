package lightGin

import "strings"

//定义好前缀树节点
type node struct{
	pattern string //待匹配的url
	part  string //当前路由中的一部分
	children []*node //子节点
	isWild bool //是否精确匹配
}

//第一个匹配成功的节点
func (n *node) matchChild(part string) *node{
	for _,child := range n.children{
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

//查询所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node{
	nodes := make([]*node,0)
	for _,child := range n.children {
		if child.part == part || child.isWild{
			nodes = append(nodes,child)
		}
	}
	return nodes
}

//向前缀树中插入
func (n *node) insert(pattern string,parts []string,height int){
	//注意只在最后一层放pattern，以避免go/:lang之类未完全匹配的返回
	if len(parts) == height{
		n.pattern = pattern
		return
	}
	part := parts[height]
	//先从child中寻找一下是否存在
	child := n.matchChild(part)
	//在children中匹配不下去了
	if child == nil{
		child = &node{part:part,isWild:part[0] == ':' || part[0]=='*'}
		n.children = append(n.children,child)
	}
	//递归的向下一层
	child.insert(pattern,parts,height+1)
}

//从前缀树中查找路径
func (n *node) search(parts []string,height int)*node{
	//退出条件是查找到最后一层或者*匹配
	if len(parts) == height || strings.HasPrefix(n.part,"*"){
		//通过pattern来进行比较
		if n.pattern == ""{
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _,child := range children{
		result := child.search(parts,height+1)
		if result != nil{
			return result
		}
	}

	return nil
}



