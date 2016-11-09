package main

import (
	"./conf"
	"./handler"
	"./service"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	conf.ParseEnvConfig()
	service.Register()

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.HandleFunc("/ping", handler.Ping)
	mux.HandleFunc("/images", handler.ListImages)
	mux.HandleFunc("/services", handler.ListServices)
	mux.HandleFunc("/apps", handler.ListApps)
	http.HandleFunc("/", handler.Warp(mux))

	log.Info("server is Listerning at 0.0.0.0:8080...")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
