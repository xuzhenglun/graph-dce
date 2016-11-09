package service

import (
	"encoding/base64"
	"net/http"

	"./../conf"

	log "github.com/Sirupsen/logrus"
	docker "github.com/docker/docker/client"
)

type DceClient struct {
	Host     string
	Username string
	Password string
	*docker.Client
}

var DefaultDceClinet *DceClient

func Register() {
	c := conf.GetConf()
	DefaultDceClinet = &DceClient{
		Host:     c.DceHost,
		Username: c.Username,
		Password: c.Password,
	}

	basic := base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))
	header := map[string]string{
		"Authorization": "Basic " + basic,
	}

	var err error
	DefaultDceClinet.Client, err = docker.NewClient(c.DceHost, "", nil, header)
	if err != nil {
		log.Error(err)
		return
	}
}

func AddHeader(req *http.Request) {
	c := conf.GetConf()

	basic := base64.StdEncoding.EncodeToString([]byte(c.Username + ":" + c.Password))
	req.Header.Add("Authorization", "Basic "+basic)
}
