package api

const DefaultServerName = "defaultServer"

//func EnforceJSONHandlerMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
//		writer.Header().Set("Content-Type", "application/json")
//		next.ServeHTTP(writer, req)
//	})
//}
