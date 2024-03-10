package router

import (
	"fmt"
	"net/http"
)

func RegisterRouter() *http.ServeMux {
	r := http.NewServeMux()
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	r.HandleFunc("GET /test1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello test1")
	})

	r.HandleFunc("/test2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello test2")
	})

	r.HandleFunc("GET /auth", HelloWorldAuth)

	r.HandleFunc("GET /auth/{provider}", beginAuthProviderCallback)

	return r
}