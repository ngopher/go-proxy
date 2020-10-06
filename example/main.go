package main

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"

	"proxy"
)

func main() {
	// gorilla/mux router
	r := mux.NewRouter()

	// context which has the url value
	ctx := context.WithValue(context.Background(), "url", &url.URL{
		Scheme: "https",
		Host:   "10.10.1.25:4999",
	})

	//
	r.Methods(http.MethodGet).
		Path("/").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(`Hello world`))
		})

	r.Methods(http.MethodPost, http.MethodOptions, http.MethodPut, http.MethodDelete, http.MethodGet).
		Path("/forward/{rest:.*}").Handler(proxy.ProxyHandler(ctx))
}
