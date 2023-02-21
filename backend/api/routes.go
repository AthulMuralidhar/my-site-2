package api

import (
	"encoding/json"
	"github.com/AthulMuralidhar/my-site-2/server"
	"log"
	"net/http"
)

func GetIndex(writer http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Printf("%s: got / request\n", ctx.Value(server.DefaultServerName))
	data := "Hello World - index page"
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(data)
}
