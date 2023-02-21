package api

import (
	"log"
	"net/http"
)

func GetIndex(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Printf("%s: got / request\n", ctx.Value(DefaultServerName))
	writer.Write([]byte("Hello World - index page"))
}
