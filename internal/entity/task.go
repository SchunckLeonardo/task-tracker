package entity

import (
	"errors"
	"github.com/SchunckLeonardo/task-tracker/pkg/utils"
	"time"
)

type ReadTaskFile struct {
	Tasks []Task `json:"tasks"`
}

const (
	ToDo       = "todo"
	InProgress = "in-progress"
	Done       = "done"
	filename   = "tasks.json"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(description string) (*Task, error) {
	tasks, err := utils.ReadFile[ReadTaskFile](filename)
	if err != nil {
		return nil, err
	}

	lastIdAddedInTask := tasks.Tasks[len(tasks.Tasks)-1].ID

	return &Task{
		ID:          lastIdAddedInTask + 1,
		Description: description,
		Status:      ToDo,
		CreatedAt:   time.Now(),
	}, nil
}

func (t *Task) Add() error {
	tasks, err := utils.ReadFile[ReadTaskFile](filename)
	if err != nil {
		return err
	}

	tasks.Tasks = append(tasks.Tasks, *t)

	err = utils.UpdateFile(filename, tasks.Tasks)
	return err
}

func (t *Task) Update(id int, newDescription string) error {
	tasks, err := utils.ReadFile[ReadTaskFile](filename)
	if err != nil {
		return err
	}

	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks[i].Description = newDescription
		}
	}

	err = utils.UpdateFile(filename, tasks.Tasks)
	return err
}

func (t *Task) Delete(id int) error {
	tasks, err := utils.ReadFile[ReadTaskFile](filename)
	if err != nil {
		return err
	}

	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
		}
	}

	err = utils.UpdateFile(filename, tasks.Tasks)
	return err
}

func (t *Task) MarkNewStatus(id int, status string) error {
	if status != InProgress && status != Done {
		return errors.New("status should be in-progress or done")
	}

	tasks, err := utils.ReadFile[ReadTaskFile](filename)
	if err != nil {
		return err
	}

	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks[i].Status = status
		}
	}

	err = utils.UpdateFile(filename, tasks.Tasks)
	return err
}

func (t *Task) ListAllTasks() ([]Task, error) {
	var allTasks []Task

	tasks, err := utils.ReadFile[ReadTaskFile](filename)
	if err != nil {
		return nil, err
	}

	for _, task := range tasks.Tasks {
		allTasks = append(allTasks, task)
	}

	return allTasks, nil
}

func (t *Task) ListTasksFilteredByStatus(status string) ([]Task, error) {
	if status != ToDo && status != InProgress && status != Done {
		return nil, errors.New("status should be todo or in-progress or done")
	}

	var allTasks []Task

	tasks, err := utils.ReadFile[ReadTaskFile](filename)
	if err != nil {
		return nil, err
	}

	for _, task := range tasks.Tasks {
		if task.Status == status {
			allTasks = append(allTasks, task)
		}
	}

	return allTasks, nil
}
