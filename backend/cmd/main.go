package main

import (
	"github.com/AthulMuralidhar/my-site-2/pkg"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", pkg.GetIndex)
}
