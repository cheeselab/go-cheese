# go-cheese
Common packages for our Golang projects

#### `router`
Is a wrapper of common HTTP router written in Golang. Currently, we support gin & mux. Below is an example with Mux router.

```go
func main() {
    // You can replace router.GIN with router.MUX
    // and change handlerGin to handlerMux
    r, err := router.New(router.GIN)
    if err != nil {
    	log.Fatal(err.Error())
    }
    r.Get("/", handlerGin)
    // r.Get("/", handlerMux)
    if err := r.Run(); err != nil {
        log.Fatal(err.Error())
    }
}

func handlerMux(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello World!"))
}

func handlerGin(c *gin.Context) {
    c.String(http.StatusOK, "Hello World!")
}
```
Middleware
```go
func main() {
    ...
    r.UseMiddleware(loggingMiddleware)
    if err := r.Run(); err != nil {
        log.Fatal(err.Error())
    }
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.RequestURI)
        next.ServeHTTP(w, r)
    })
}
```