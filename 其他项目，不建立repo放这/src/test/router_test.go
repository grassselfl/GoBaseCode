package test

import (
	"fmt"
	"hammer/src/core"
	"net/http"
	"testing"
)

func Test(t *testing.T) {
	testRoutes := []struct {
		method     string
		path       string
		handleFunc core.HandleFunc
	}{
		{method: http.MethodGet, path: "/", handleFunc: func(ctx *core.Context) {}},
		{method: http.MethodGet, path: "//user", handleFunc: func(ctx *core.Context) {}},
		{method: http.MethodGet, path: "user", handleFunc: func(ctx *core.Context) {}},
		{method: http.MethodGet, path: "user/", handleFunc: func(ctx *core.Context) {}},
		{method: http.MethodGet, path: "/user", handleFunc: func(ctx *core.Context) {}},
		{method: http.MethodGet, path: "/user/home", handleFunc: func(ctx *core.Context) {}},
		{method: http.MethodGet, path: "/order/detail", handleFunc: func(ctx *core.Context) {}},
		{method: http.MethodGet, path: "/order/create", handleFunc: func(ctx *core.Context) {}},
		{method: http.MethodGet, path: "/login", handleFunc: func(ctx *core.Context) {}},
	}

	r := core.NewRouter()
	for _, tr := range testRoutes {
		r.AddRoute(tr.method, tr.path, tr.handleFunc)
	}
	fmt.Println("123")
}
