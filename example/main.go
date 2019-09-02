package main

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"

	"github.com/johansetia/jowt"
)

type (
	// Claims :
	Claims struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}

	// TokenJWT is used to verify create JWT
	TokenJWT struct {
		Iss    string `json:"iss"`
		UID    string `json:"uid"`
		Aud    string `json:"aud"`
		Claims `json:"claims"`
	}

	// CreatedToken is used to store token that has been created by Token Struct
	CreatedToken struct {
		// Token string
		Token string `json:"token"`
	}

	// CostomMux is used to store mux from http.Servemux Struct and your costom middlewares
	CostomMux struct {
		http.ServeMux
		middlewares []func(next http.Handler) http.Handler
	}
)

var mt TokenJWT
var key = "qwertyuiopasdfghjklzxcvbnm123456"
var m = new(jowt.Security)

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

func main() {
	m.SecretKey = key
	m.WhiteListURI = []string{"/auth"}
	m.Message = map[string]interface{}{
		"rc":      500,
		"message": "Error Authentication",
	}

	mux := new(CostomMux)
	mux.RegisterMiddleware(m.JWTMiddleware)
	mux.HandleFunc("/get-server-info", getInfo)
	mux.HandleFunc("/auth", auth)
	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  time.Second * (1 / 2),
		WriteTimeout: time.Second * (1 / 2),
	}

	server.ListenAndServe()
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	serverInfo := map[string]interface{}{
		"OS":     runtime.GOOS,
		"Thread": runtime.NumCPU(),
	}
	if !claimsValidator() {
		isError(w)
		return
	}

	jsonString, _ := json.Marshal(serverInfo)
	w.Header().Set("Content-Type", "Application/json")
	w.Write(jsonString)
}

func auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		isError(w)
	}
	decoder := json.NewDecoder(r.Body)
	var wilToken jowt.Payload
	decoder.Decode(&wilToken)

	token, ok := jowt.HS512(key).SetPayload(wilToken).Get()
	if !ok {
		isError(w)
	}
	var s CreatedToken

	s.Token = token
	send, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "Application/json")
	w.Write(send)
}

func claimsValidator() bool {
	byteToken, err := json.Marshal(m.MiddlewarePayload)
	if err != nil {
		return false
	}
	err = json.Unmarshal(byteToken, &mt)
	if err != nil {
		return false
	}
	if mt.Aud == "" || mt.Email == "" || mt.UID == "" || mt.Username == "" {
		return false
	}
	return true
}

func isError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "Application/json")
	w.Write([]byte(`{"status":"can't read your request body."}`))
}
