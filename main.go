package main

import (
	"./conf"
	"./handler"
	"./middleware"
	"./router"
	"./service"

	"net/http"

	log "github.com/Sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	conf.ParseEnvConfig()
	service.Register()

	mux := router.NewRouter()
	mux.Use("/ping", handler.Ping)
	mux.Use("/", handler.Static)

	api := mux.NewSubRouter()
	api.UseFunc(middleware.Warp)
	api.Use("/images", handler.ListImages)
	api.Use("/services", handler.ListServices)
	api.Use("/apps", handler.ListApps)
	mux.Register()

	log.Info("server is Listerning at 0.0.0.0:8080...")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
