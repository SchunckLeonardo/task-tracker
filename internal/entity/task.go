package entity

import (
	"errors"
	"fmt"
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

func NewTask() *Task {
	return &Task{
		Status:    ToDo,
		CreatedAt: time.Now(),
	}
}

func (t *Task) Add(description string) error {
	tasks, err := utils.ReadFile[ReadTaskFile](filename)
	if err != nil {
		return err
	}

	var lastIdAddedInTask int
	if len(tasks.Tasks) < 1 {
		lastIdAddedInTask = 1
	} else {
		lastIdAddedInTask = tasks.Tasks[len(tasks.Tasks)-1].ID + 1
	}

	t.Description = description
	t.ID = lastIdAddedInTask

	tasks.Tasks = append(tasks.Tasks, *t)

	err = utils.UpdateFile(filename, tasks.Tasks)
	return err
}

func (t *Task) Update(id int, newDescription string) error {
	tasks, err := utils.ReadFile[ReadTaskFile](filename)
	if err != nil {
		return err
	}

	founded := false

	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks[i].Description = newDescription
			tasks.Tasks[i].UpdatedAt = time.Now()
			founded = true
			break
		}
	}

	if founded == false {
		return fmt.Errorf("not found task with id: %v", id)
	}

	err = utils.UpdateFile(filename, tasks.Tasks)
	return err
}

func (t *Task) Delete(id int) error {
	tasks, err := utils.ReadFile[ReadTaskFile](filename)
	if err != nil {
		return err
	}

	founded := false

	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i+1:]...)
			founded = true
			break
		}
	}

	if founded == false {
		return fmt.Errorf("not found task with id: %v", id)
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

	founded := false

	for i := range tasks.Tasks {
		if tasks.Tasks[i].ID == id {
			tasks.Tasks[i].Status = status
			founded = true
			break
		}
	}

	if founded == false {
		return fmt.Errorf("not found task with id: %v", id)
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

	if len(tasks.Tasks) < 1 {
		return nil, fmt.Errorf("need to add a new task to list tasks")
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

	if len(tasks.Tasks) < 1 {
		return nil, fmt.Errorf("need to add a new task to list tasks")
	}

	for _, task := range tasks.Tasks {
		if task.Status == status {
			allTasks = append(allTasks, task)
		}
	}

	return allTasks, nil
}
