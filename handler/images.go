package handler

import (
	"encoding/json"
	"net/http"

	"../service"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"strings"
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

	imageMap := make(map[string][]string)
	for _, image := range images {
		for _, tag := range image.RepoTags {
			s := strings.SplitN(tag, ":", 2)
			if len(s) != 2 {
				continue
			}

			imageMap[s[0]] = append(imageMap[s[0]], tag)
		}
	}

	result := make([]*Image, 0, len(imageMap))
	for image, tags := range imageMap {
		i := Image{
			Name: image,
			Tags: tags,
		}
		result = append(result, &i)
	}

	json.NewEncoder(w).Encode(result)
}
