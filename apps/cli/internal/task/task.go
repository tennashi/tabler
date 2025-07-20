package task

import "time"

type Task struct {
	ID        string
	Title     string
	Deadline  time.Time
	Priority  int
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTask(id, title string, deadline time.Time, priority int) *Task {
	now := time.Now()
	return &Task{
		ID:        id,
		Title:     title,
		Deadline:  deadline,
		Priority:  priority,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type TaskTag struct {
	TaskID string
	Tag    string
}

func NewTaskTag(taskID, tag string) *TaskTag {
	return &TaskTag{
		TaskID: taskID,
		Tag:    tag,
	}
}
