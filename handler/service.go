package handler

import (
	"./../service"

	"encoding/json"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

type Service struct {
	Name     string `json:"name"`
	Image    string `json:"image"`
	App      string `json:"app"`
	Running  int    `json:"running"`
	Replicas *int   `json:"replicas"`
}

func ListServices(w http.ResponseWriter, r *http.Request) {
	services, err := service.DefaultDceClinet.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		log.Error(err)
		return
	}

	result := make([]*Service, 0, len(services))
	for _, service := range services {
		images := service.Spec.TaskTemplate.ContainerSpec.Image
		array := strings.Split(images, ":")
		if len(array) == 1 {
			array = append(array, "latest")
		}

		s := Service{
			Name:  service.Spec.Name,
			Image: strings.Join(array, ":"),
		}

		appName, _ := service.Spec.Labels["com.docker.stack.namespace"]
		s.App = appName
		if service.Spec.Mode.Replicated != nil {
			r := int(*service.Spec.Mode.Replicated.Replicas)
			s.Replicas = &r
		}

		if _, ok := service.Spec.Labels["io.daocloud.dce.system"]; !ok {
			s.Running = GetStatusByServiceName(service.Spec.Name)
			result = append(result, &s)
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}
