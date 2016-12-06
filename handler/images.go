package handler

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"../service"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

type Image struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type Images []*Image

func (s Images) Len() int           { return len(s) }
func (s Images) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Images) Less(i, j int) bool { return s[i].Name < s[j].Name }

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

	result := make(Images, 0, len(imageMap))
	for image, tags := range imageMap {
		i := Image{
			Name: image,
			Tags: tags,
		}
		result = append(result, &i)
	}

	sort.Sort(result)
	json.NewEncoder(w).Encode(result)
}
