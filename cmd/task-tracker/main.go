package main

import (
	"fmt"
	"github.com/SchunckLeonardo/task-tracker/internal/entity"
	"log"
	"os"
	"strconv"
	"strings"
)

const filename = "tasks.json"

func init() {
	file, err := os.ReadFile(filename)
	if err != nil {
		_, err = os.Create(filename)
		if err != nil {
			panic("error to create tasks.json")
		}
	}

	if !strings.Contains(string(file), "tasks") {
		err = os.WriteFile(filename, []byte(`{"tasks": []}`), 0666)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	task := entity.NewTask()

	switch os.Args[1] {
	case "add":
		err := task.Add(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Task added successfully (ID: %d)", task.ID)
	case "update":
		idParsedToInt, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		err = task.Update(idParsedToInt, os.Args[3])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Task updated successfully")
	case "delete":
		idParsedToInt, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		err = task.Delete(idParsedToInt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Task deleted successfully")
	case "mark-in-progress":
		idParsedToInt, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		err = task.MarkNewStatus(idParsedToInt, entity.InProgress)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Task marked in progress successfully")
	case "mark-done":
		idParsedToInt, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		err = task.MarkNewStatus(idParsedToInt, entity.Done)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Task marked done successfully")
	case "list":
		var filter string
		if len(os.Args) > 2 {
			filter = os.Args[2]
		}

		if filter == "todo" {
			tasks, err := task.ListTasksFilteredByStatus(entity.ToDo)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%#v", tasks)
		} else if filter == "in-progress" {
			tasks, err := task.ListTasksFilteredByStatus(entity.InProgress)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%#v", tasks)
		} else if filter == "done" {
			tasks, err := task.ListTasksFilteredByStatus(entity.Done)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%#v", tasks)
		} else {
			tasks, err := task.ListAllTasks()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%#v", tasks)
		}
	}
}
