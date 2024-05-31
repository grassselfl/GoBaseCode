package core

import "net/http"

// Server 接口
type Server interface {
	http.Handler
	Start(address string) error
	AddRoute(method string, path string, handler HandleFunc)
}
