package server

import (
	"net/http"
	"time"
)

type Handler struct {
	Path string
	Func http.HandlerFunc
}

func NewServer(port string, handlers []Handler, middleware ...func(http.Handler) http.Handler) func() {
	mux := http.NewServeMux()
	for _, h := range handlers {
		mux.Handle(h.Path, h.Func)
	}

	var finalHandler http.Handler = mux
	for i := len(middleware) - 1; i >= 0; i-- { // Reverse order for correct chaining
		finalHandler = middleware[i](finalHandler)
	}

	srv := &http.Server{
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  5 * time.Second,
		Handler:      finalHandler,
	}
	return func() {
		panic(srv.ListenAndServe())
	}
}
