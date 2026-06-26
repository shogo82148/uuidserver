package main

import (
	_ "embed" // for go:embed
	"net/http"
	"strconv"
	"uuid"

	"github.com/shogo82148/ridgenative"
)

//go:embed favicon.ico
var favicon []byte

type UUIDHandlerFunc func() uuid.UUID

// ServeHTTP implements http.Handler interface.
// based on https://github.com/syumai/uuidserver/blob/7a24e20453174f88c5d762356c9a14cad3cd2711/uuidserver/server.go
func (g UUIDHandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	b, _ := g().MarshalText()
	contentLength := len(b)
	w.Header().Set("Content-Length", strconv.Itoa(contentLength))
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.Write(b)
}

func faviconHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Length", strconv.Itoa(len(favicon)))
	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", "public, max-age=604800") // 7 days
	w.Write(favicon)
}

func main() {
	http.HandleFunc("GET /favicon.ico", faviconHandler)
	http.Handle("GET /", UUIDHandlerFunc(uuid.New))
	http.Handle("GET /v4", UUIDHandlerFunc(uuid.NewV4))
	http.Handle("GET /v7", UUIDHandlerFunc(uuid.NewV7))
	http.Handle("GET /nil", UUIDHandlerFunc(uuid.Nil))
	http.Handle("GET /max", UUIDHandlerFunc(uuid.Max))
	ridgenative.ListenAndServe(":8080", nil)
}
