package handler

import (
	"../service"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"golang.org/x/net/context"
)

type Tasks struct {
	t []swarm.Task
}

func GetAllTasks() Tasks {
	option := types.TaskListOptions{}
	tasks, err := service.DefaultDceClinet.TaskList(context.Background(), option)
	if err != nil {
		log.Error(err)
		return Tasks{t: []swarm.Task{}}
	}

	return Tasks{t: tasks}
}

func (tasks Tasks) GetStatusByServiceId(id string) (running int) {
	for _, task := range tasks.t {
		if task.ServiceID == id && task.Status.State == swarm.TaskStateRunning {
			running++
		}
	}
	return
}
