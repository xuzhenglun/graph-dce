package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"./../service"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

type Service struct {
	Id       string `json:"id"`
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

	tasks := GetAllTasks()

	result := make([]*Service, 0, len(services))
	for _, service := range services {
		images := service.Spec.TaskTemplate.ContainerSpec.Image
		array := strings.Split(images, ":")
		if len(array) == 1 {
			array = append(array, "latest")
		}

		s := Service{
			Id:    service.ID,
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
			s.Running = tasks.GetStatusByServiceId(service.ID)
			result = append(result, &s)
		}
	}

	json.NewEncoder(w).Encode(result)
}
