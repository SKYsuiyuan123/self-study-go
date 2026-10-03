package util

import (
	"fmt"
	"time"
)

// 任务
type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

func PrintTasks(list []Task) {
	for _, val := range list {
		fmt.Printf("#%d %s \n", val.Id, val.Title)
	}
}
