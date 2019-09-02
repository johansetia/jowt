package jowt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/johansetia/jowt/helper"
)

// Jwt is a middleware to verify token that you set before.
func Jwt(m *Security, next http.Handler) http.Handler {
	jwt := func(w http.ResponseWriter, r *http.Request) {
		uri := strings.Split(fmt.Sprintf("%s", r.URL), "?")[0]
		found, _ := helper.InArray(uri, m.WhiteListURI)
		if !found {
			h := r.Header.Get("Authorization")
			if h == "" {
				errorAuth(m, w)
				return
			}
			token := strings.Split(h, "Bearer ")[1]
			if token == "" {
				errorAuth(m, w)
				return
			}
			authorization := Verify(m.SecretKey).SetToken(token)
			if !authorization.Status() {
				errorAuth(m, w)
				return
			}
			m.SetPayloadFromMiddleware(authorization.Payload)
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(jwt)
}

func errorAuth(m *Security, w http.ResponseWriter) {
	error, _ := json.Marshal(m.Message)
	w.Header().Set("Content-Type", "Application/json")
	w.Write(error)
}
