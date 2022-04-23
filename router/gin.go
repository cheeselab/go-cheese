package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	RouterGin struct {
		routes []RouteGin
	}

	RouteGin struct {
		Method  string
		Path    string
		Handler func(*gin.Context)
	}
)

func NewginRouter() *RouterGin {
	return &RouterGin{}
}

func (r *RouterGin) addRoute(method, path string, handleFunc interface{}) {
	r.routes = append(r.routes, RouteGin{
		Method:  method,
		Path:    path,
		Handler: handleFunc.(func(*gin.Context)),
	})
}

func (r *RouterGin) Get(path string, handleFunc interface{}) {
	r.addRoute(http.MethodGet, path, handleFunc)
}

func (r *RouterGin) Post(path string, handleFunc interface{}) {
	r.addRoute(http.MethodPost, path, handleFunc)
}

func (r *RouterGin) Put(path string, handleFunc interface{}) {
	r.addRoute(http.MethodPut, path, handleFunc)
}

func (r *RouterGin) Patch(path string, handleFunc interface{}) {
	r.addRoute(http.MethodPatch, path, handleFunc)
}

func (r *RouterGin) Delete(path string, handleFunc interface{}) {
	r.addRoute(http.MethodDelete, path, handleFunc)
}

func (r *RouterGin) Options(path string, handleFunc interface{}) {
	r.addRoute(http.MethodOptions, path, handleFunc)
}

func (r *RouterGin) Run() error {
	router := gin.Default()
	for _, route := range r.routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Path, route.Handler)
		case http.MethodPost:
			router.POST(route.Path, route.Handler)
		case http.MethodPut:
			router.PUT(route.Path, route.Handler)
		case http.MethodPatch:
			router.PATCH(route.Path, route.Handler)
		case http.MethodDelete:
			router.DELETE(route.Path, route.Handler)
		case http.MethodOptions:
			router.OPTIONS(route.Path, route.Handler)
		}
	}
	return router.Run(":8000")
}
