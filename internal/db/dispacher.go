package db

import (
	"encoding/json"
	"fmt"
)

type Parametros struct {
	Object  string
	Command string
	Args    []byte
}

func NewParametros(object, command string, args []byte) *Parametros {
	return &Parametros{
		Object:  object,
		Command: command,
		Args:    args,
	}
}

type Dispacher struct {
	skeleton *Skeleton
}

func NewDispacher(skeleton *Skeleton) *Dispacher {
	return &Dispacher{skeleton: skeleton}
}

func (d *Dispacher) Invoke(p *Parametros) ([]byte, error) {
	switch p.Object {
	case "Task":
		switch p.Command {
		case "InsertTask":
			return d.skeleton.InsertTask(p.Args)
		case "GetAllTasks":
			return d.skeleton.GetAllTasks()
		case "GetTaskByID":
			return d.skeleton.GetTaskByID(p.Args)
		case "DeleteTask":
			return d.skeleton.DeleteTask(p.Args)
		default:
			return json.Marshal(fmt.Sprintf("comando nao reconhecido: %s", p.Command))
		}
	default:
		return json.Marshal(fmt.Sprintf("objeto nao reconhecido: %s", p.Object))
	}
}
