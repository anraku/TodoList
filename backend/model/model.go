package model

import "github.com/anraku/TodoList/backend/helper"

type Task struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func (task *Task) ToResponseMap() helper.ResponseMap {
	return helper.ResponseMap{
		"id":   task.ID,
		"text": task.Text,
		"done": task.Done,
	}
}
