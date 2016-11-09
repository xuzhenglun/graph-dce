package handler

import (
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong!"))
}

func Static(w http.ResponseWriter, r *http.Request) {
	h := http.FileServer(http.Dir("./static"))
	h.ServeHTTP(w, r)
}
