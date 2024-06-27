package server

import (
	"net/http"
	"time"
)

type Handler struct {
	Path string
	Func http.HandlerFunc
}

func NewServer(port string, handlers ...Handler) func() {
	mux := http.NewServeMux()
	for _, h := range handlers {
		mux.Handle(h.Path, h.Func)
	}
	srv := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  5 * time.Second,
		Handler:      mux,
	}
	return func() {
		panic(srv.ListenAndServe())
	}
}
