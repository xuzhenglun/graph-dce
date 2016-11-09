package handler

import (
	"../service"

	"encoding/json"
	"golang.org/x/net/context"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types"
)

type Image struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func ListImages(w http.ResponseWriter, r *http.Request) {
	images, err := service.DefaultDceClinet.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		log.Error()
		return
	}

	result := make([]*Image, 0, len(images))
	for _, image := range images {
		for _, tag := range image.RepoTags {
			i := Image{
				Name: tag,
			}
			result = append(result, &i)
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}
