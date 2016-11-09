package handler

import (
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong!"))
}

func Option(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	w.Header().Set("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static", 301)
}
