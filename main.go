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

	//router := mux.NewRouter()
	http.HandleFunc("/", handler.Index)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/ping", handler.Ping)
	http.HandleFunc("/images", handler.ListImages)
	http.HandleFunc("/services", handler.ListServices)
	http.HandleFunc("/apps", handler.ListApps)
	http.HandleFunc("/{.*}", handler.Option)

	log.Info("server is Listerning at 0.0.0.0:8080...")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
