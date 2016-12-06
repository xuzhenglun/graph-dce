package handler

import (
	"encoding/json"
	"net/http"
	"sort"
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

type Services []*Service

func (s Services) Len() int           { return len(s) }
func (s Services) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Services) Less(i, j int) bool { return s[i].Name < s[j].Name }

func ListServices(w http.ResponseWriter, r *http.Request) {
	services, err := service.DefaultDceClinet.ServiceList(context.Background(), types.ServiceListOptions{})
	if err != nil {
		log.Error(err)
		return
	}

	tasks := GetAllTasks()

	result := make(Services, 0, len(services))
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

	sort.Sort(result)
	json.NewEncoder(w).Encode(result)
}
