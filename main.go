package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

// main
func main() {
	http.Handle("/", loggingMiddleware(http.HandlerFunc(handler)))
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Package main #HttpServerGo\n")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("uri: %s", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
