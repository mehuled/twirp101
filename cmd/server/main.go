package main

import (
	"net/http"

	"github.com/mehuled/twirp/haberdasher"
	"github.com/mehuled/twirp/internal/haberdasherserver"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
