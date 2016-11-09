package handler

import (
	"../service"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
)

type Task struct {
	Name   string
	Status string
}

func ListTasksByServiceName(s string) []swarm.Task {
	f := filters.NewArgs()
	f.Add("service", s)
	option := types.TaskListOptions{
		Filters: f,
	}
	tasks, err := service.DefaultDceClinet.TaskList(context.Background(), option)
	if err != nil {
		log.Error(err)
		return nil
	}

	return tasks
}

func GetStatusByServiceName(s string) (running int) {
	tasks := ListTasksByServiceName(s)
	for _, task := range tasks {
		if task.Status.State == swarm.TaskStateRunning {
			running++
		}
	}
	return
}
