package main

type Task struct {
	ID   int
	Name string
	Done bool
}

func ListTasks() []Task {
	tasks := []Task{
		{
			ID:   1,
			Name: "testing 1",
			Done: false},
		{
			ID:   2,
			Name: "testing 2",
			Done: false},
		{
			ID:   3,
			Name: "testing 3",
			Done: true},
	}
	return tasks
}
