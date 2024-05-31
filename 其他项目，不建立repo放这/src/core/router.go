package core

import "strings"

type router struct {
	routerNodes map[string]*routerNode
}

type routerNode struct {
	path        string
	routerNodes map[string]*routerNode
	handlerFunc HandleFunc
}

func (r routerNode) getOrCreateNode(path string) *routerNode {
	if r.routerNodes == nil {
		r.routerNodes = make(map[string]*routerNode)
	}
	child, ok := r.routerNodes[path]
	if !ok {
		child = &routerNode{path: path}
		r.routerNodes[path] = child
	}
	return child
}

func NewRouter() *router {
	return &router{
		routerNodes: map[string]*routerNode{},
	}
}

func (r *router) AddRoute(method string, path string, handleFunc HandleFunc) {

	// 获取方法子路由，没有自动创建
	root, ok := r.routerNodes[method]
	if !ok {
		root = &routerNode{path: "/"}
		r.routerNodes[method] = root
	}

	segments := strings.Split(path, "/")
	for _, s := range segments {
		root = root.getOrCreateNode(s)
	}

	root.handlerFunc = handleFunc

}
