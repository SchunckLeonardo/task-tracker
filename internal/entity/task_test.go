package entity_test

import (
	"github.com/SchunckLeonardo/task-tracker/internal/entity"
	"github.com/SchunckLeonardo/task-tracker/pkg/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const fileName = "tasks.json"

func Setup(t *testing.T) {
	json := `{"tasks": [{"id": 1, "status": "todo"}, {"id": 2, 
"status": "todo"}]}`

	_, err := os.ReadFile(fileName)
	if err != nil {
		_, err = os.Create(fileName)
		assert.Nil(t, err)
	}

	err = os.WriteFile(fileName, []byte(json), 0666)
	assert.Nil(t, err)
}

func TearDown(t *testing.T) {
	err := os.Remove(fileName)
	assert.Nil(t, err)
}

func TestTask_Add(t *testing.T) {
	Setup(t)
	defer TearDown(t)

	task := entity.NewTask()

	err := task.Add("Buy new book")
	assert.Nil(t, err)

	tasks, err := utils.ReadFile[entity.ReadTaskFile](fileName)
	assert.Nil(t, err)

	assert.Equal(t, tasks.Tasks[2].ID, 3)
	assert.Equal(t, tasks.Tasks[2].Description, "Buy new book")
	assert.Equal(t, tasks.Tasks[2].Status, entity.ToDo)
}

func TestTask_Update(t *testing.T) {
	Setup(t)
	defer TearDown(t)

	task := entity.NewTask()

	err := task.Add("Buy new book")
	assert.Nil(t, err)

	err = task.Update(3, "Go shopping")
	assert.Nil(t, err)

	tasks, err := utils.ReadFile[entity.ReadTaskFile](fileName)
	assert.Nil(t, err)

	assert.Equal(t, tasks.Tasks[2].ID, 3)
	assert.Equal(t, tasks.Tasks[2].Description, "Go shopping")
	assert.Equal(t, tasks.Tasks[2].Status, entity.ToDo)
}

func TestTask_Delete(t *testing.T) {
	Setup(t)
	defer TearDown(t)

	task := entity.NewTask()

	err := task.Add("Buy new book")
	assert.Nil(t, err)

	tasks, err := utils.ReadFile[entity.ReadTaskFile](fileName)
	assert.Nil(t, err)

	assert.Equal(t, len(tasks.Tasks), 3)

	err = task.Delete(3)
	assert.Nil(t, err)

	tasks, err = utils.ReadFile[entity.ReadTaskFile](fileName)
	assert.Nil(t, err)

	assert.Equal(t, len(tasks.Tasks), 2)
}

func TestTask_MarkNewStatus(t *testing.T) {
	Setup(t)
	defer TearDown(t)

	task := entity.NewTask()

	err := task.Add("Buy new book")
	assert.Nil(t, err)

	err = task.MarkNewStatus(3, entity.InProgress)
	assert.Nil(t, err)

	tasks, err := utils.ReadFile[entity.ReadTaskFile](fileName)
	assert.Nil(t, err)

	assert.Equal(t, tasks.Tasks[2].Status, entity.InProgress)

	err = task.MarkNewStatus(3, entity.Done)
	assert.Nil(t, err)

	tasks, err = utils.ReadFile[entity.ReadTaskFile](fileName)
	assert.Nil(t, err)

	assert.Equal(t, tasks.Tasks[2].Status, entity.Done)
}

func TestTask_ListAllTasks(t *testing.T) {
	Setup(t)
	defer TearDown(t)

	task := entity.NewTask()

	err := task.Add("Buy new book")
	assert.Nil(t, err)

	err = task.MarkNewStatus(3, entity.InProgress)
	assert.Nil(t, err)

	task = entity.NewTask()

	err = task.Add("Go shopping")
	assert.Nil(t, err)

	task = entity.NewTask()

	err = task.Add("Play with my friends Roblox")
	assert.Nil(t, err)

	allTasks, err := task.ListAllTasks()
	assert.Nil(t, err)

	assert.Equal(t, len(allTasks), 5)
	assert.Equal(t, allTasks[len(allTasks)-1].Description, "Play with my friends Roblox")
}

func TestTask_ListTasksFilteredByStatus(t *testing.T) {
	Setup(t)
	defer TearDown(t)

	task := entity.NewTask()

	err := task.Add("Buy new book")
	assert.Nil(t, err)

	err = task.MarkNewStatus(3, entity.InProgress)
	assert.Nil(t, err)

	task = entity.NewTask()

	err = task.Add("Go shopping")
	assert.Nil(t, err)

	task = entity.NewTask()

	err = task.Add("Play with my friends Roblox")
	assert.Nil(t, err)

	err = task.MarkNewStatus(task.ID, entity.Done)
	assert.Nil(t, err)

	allTasks, err := task.ListTasksFilteredByStatus(entity.ToDo)
	assert.Nil(t, err)

	assert.Equal(t, len(allTasks), 3)

	allTasks, err = task.ListTasksFilteredByStatus(entity.InProgress)
	assert.Nil(t, err)

	assert.Equal(t, len(allTasks), 1)

	allTasks, err = task.ListTasksFilteredByStatus(entity.Done)
	assert.Nil(t, err)

	assert.Equal(t, len(allTasks), 1)
}
