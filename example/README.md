# Example of the JOWT Implementation

### 1. Make A JOWT Middleware
Set it into the global variable.
```go
var key = "qwertyuiopasdfghjklzxcvbnm123456"
var m = new(jowt.Security)
```

### 2. Config JOWT Middleware
Make a configuration by a variable instance from **jowt.Security**. you can fill the SecretKey using your secret key, and fill the Whitelist URI using string slice to set URI that won't use the JOWT Middleware and fill the Message to change an error message while validation middleware is error. you can place the config into your *func main* with your server or another.
```go
m.SecretKey = key
m.WhiteListURI = []string{"/auth"}
m.Message = map[string]interface{}{
    "rc":      500,
    "message": "Error Authentication",
}
```
### 3. Implement JOWT Middleware
there. i have a costom multiplexer, so i use that to make it easy when i will add another middleware.

```go
// CostomMux is used to store mux from http.Servemux Struct and your costom middlewares
type CostomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

// RegisterMiddleware is used to add your costom middleware.
func (c *CostomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

// ServeHTTP always be called after has a request.
func (c *CostomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux
	for _, next := range c.middlewares {
		current = next(current)
	}
	current.ServeHTTP(w, r)
}

```
and i will implement the jowt that has been configured in the step 2
```go
mux := new(CostomMux)
mux.RegisterMiddleware(m.JWTMiddleware) // Implement here
mux.HandleFunc("/get-server-info", getInfo)
mux.HandleFunc("/auth", auth)

server := http.Server{
    Addr:         ":8080",
    Handler:      mux,
    ReadTimeout:  time.Second * (1 / 2),
    WriteTimeout: time.Second * (1 / 2),
}

server.ListenAndServe()
```
### 4. Costomize your own amazing product.


made with :blue_heart: by johan setiawan 