# Task Tracker CLI
Challenge access - https://roadmap.sh/projects/task-tracker

### How to use
```
1) git clone https://github.com/SchunckLeonardo/task-tracker.git

2) go mod tidy

3) go build cmd/task-tracker/main.go

4) ./task-tracker
```

### Commands
Add new task
```shell
./task-tracker add "Buy new books"
```

Update an existing task
```shell
./task-tracker update 1 "Buy books about finance"
```

Delete an existing task
```shell
./task-tracker delete 1
```

Marking a task as in progress or done
```shell
./task-tracker mark-in-progress 1
```

```shell
./task-tracker mark-done 1
```

Listing all tasks
```shell
./task-tracker list
```

Listing tasks by status
```shell
./task-tracker list done
```

```shell
./task-tracker list todo
```

```shell
./task-tracker list in-progress
```