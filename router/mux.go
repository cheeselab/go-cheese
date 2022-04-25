package router

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type (
	RouterMux struct {
		routes      []RouteMux
		middlewares []func(http.Handler) http.Handler
	}

	RouteMux struct {
		Method  string
		Path    string
		Handler func(w http.ResponseWriter, r *http.Request)
	}
)

func NewMuxRouter() *RouterMux {
	return &RouterMux{}
}

func (r *RouterMux) addRoute(method, path string, handleFunc interface{}) {
	r.routes = append(r.routes, RouteMux{
		Method:  method,
		Path:    path,
		Handler: handleFunc.(func(http.ResponseWriter, *http.Request)),
	})
}

func (r *RouterMux) Get(path string, handleFunc interface{}) {
	r.addRoute(http.MethodGet, path, handleFunc)
}

func (r *RouterMux) Post(path string, handleFunc interface{}) {
	r.addRoute(http.MethodPost, path, handleFunc)
}

func (r *RouterMux) Put(path string, handleFunc interface{}) {
	r.addRoute(http.MethodPut, path, handleFunc)
}

func (r *RouterMux) Patch(path string, handleFunc interface{}) {
	r.addRoute(http.MethodPatch, path, handleFunc)
}

func (r *RouterMux) Delete(path string, handleFunc interface{}) {
	r.addRoute(http.MethodDelete, path, handleFunc)
}

func (r *RouterMux) Options(path string, handleFunc interface{}) {
	r.addRoute(http.MethodOptions, path, handleFunc)
}

func (r *RouterMux) Run() error {
	muxRouter := mux.NewRouter()
	// Set routes
	for _, route := range r.routes {
		muxRouter.HandleFunc(route.Path, route.Handler).Methods(strings.ToUpper(route.Method))
	}
	// Set middlewares
	for _, m := range r.middlewares {
		muxRouter.Use(m)
	}
	srv := &http.Server{
		Handler: muxRouter,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func (r *RouterMux) UseMiddleware(middleware interface{}) {
	r.middlewares = append(r.middlewares, middleware.(func(http.Handler) http.Handler))
}
