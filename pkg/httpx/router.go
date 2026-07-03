package httpx

import "net/http"

type Router interface {
	HandleFunc(pattern string, handler http.HandlerFunc)
}

type HandlerFunc func(w http.ResponseWriter, r *http.Request)
