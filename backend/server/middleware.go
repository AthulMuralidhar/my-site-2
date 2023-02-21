package server

import "net/http"

func DefaultMiddleware(handleFunc func(writer http.ResponseWriter, req *http.Request)) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		http.HandlerFunc(handleFunc).ServeHTTP(writer, req)
	})
}
