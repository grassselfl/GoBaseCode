package core

import "net/http"

type HTTPServer struct {
	*router
}

func newHtTTPServer() *HTTPServer {
	return &HTTPServer{router: NewRouter()}
}

func (s *HTTPServer) AddRoute(method string, path string, handler HandleFunc) {
	//TODO implement me
	panic("implement me")
}

func (s *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.serve(&Context{
		Request:  request,
		Response: writer,
	})
}

func (s *HTTPServer) serve(ctx *Context) {
	// 查找路由，执行命中的业务逻辑
}

func (s *HTTPServer) Start(address string) error {
	return http.ListenAndServe(address, s)
}

// 确保接口实现
var _ Server = &HTTPServer{}
