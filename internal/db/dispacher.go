package db

import (
	"encoding/json"
	"fmt"
)

type Dispacher struct {
	skeleton *Skeleton
}

func NewDispacher(skeleton *Skeleton) *Dispacher {
	return &Dispacher{skeleton: skeleton}
}

func (d *Dispacher) Invoke(object string, command string, args []byte) ([]byte, error) {
	switch object {
	case "Task":
		return d.InvokeTask(command, args)
	default:
		return json.Marshal(fmt.Sprintf("comando não reconhecido: %s", command))
	}
}

func (d *Dispacher) InvokeTask(command string, args []byte) ([]byte, error) {
	switch command {
	case "InsertTask":
		return d.skeleton.InsertTask(args)
	case "GetAllTasks":
		return d.skeleton.GetAllTasks()
	case "GetTaskByID":
		return d.skeleton.GetTaskByID(args)
	case "DeleteTask":
		return d.skeleton.DeleteTask(args)
	default:
		return json.Marshal(fmt.Sprintf("comando não reconhecido: %s", command))
	}
}

