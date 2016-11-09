package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"./../conf"
	"./../service"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types/swarm"
)

type App struct {
	Name     string     `json:"name"`
	Services []*Service `json:"services"`
}

type DceApp struct {
	Name     string
	Tenant   string
	Services []*swarm.Service
}

func ListApps(w http.ResponseWriter, r *http.Request) {
	c := conf.GetConf()
	req, err := http.NewRequest("GET", c.DceHost+"/api/apps", nil)
	if err != nil {
		log.Errorf("failed to new request: %v", err)
		return
	}
	service.AddHeader(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("failed to connect remote dce: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		w.WriteHeader(resp.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dceApps := []*DceApp{}
	if err := json.Unmarshal(body, &dceApps); err != nil {
		log.Error(err)
		log.Info(string(body))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := make([]*App, 0, len(dceApps))

	for _, dceApp := range dceApps {
		app := App{
			Name: dceApp.Name,
		}
		app.Services = make([]*Service, 0, len(dceApp.Services))
		for _, service := range dceApp.Services {
			s := Service{
				Name:  service.Spec.Name,
				Image: service.Spec.TaskTemplate.ContainerSpec.Image,
			}

			appName, _ := service.Spec.Labels["com.docker.stack.namespace"]
			s.App = appName

			if _, ok := service.Spec.Labels["io.daocloud.dce.system"]; !ok {
				app.Services = append(app.Services, &s)
			}
		}
		result = append(result, &app)
	}

	json.NewEncoder(w).Encode(result)
}
