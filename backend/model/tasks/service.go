package tasks

import (
	"github.com/anraku/TodoList/backend/model"
	"github.com/jinzhu/gorm"
)

type Model struct {
	DB *gorm.DB
}

const table = "tasks"

func New(db *gorm.DB) *Model {
	return &Model{
		DB: db,
	}
}

func (m *Model) FindAll() (tasks []model.Task, err error) {
	err = m.DB.Table(table).Find(&tasks).Error
	return
}

func (m *Model) Create(task model.Task) (t model.Task, err error) {
	t = task
	err = m.DB.Table(table).Create(&t).Error
	return
}

func (m *Model) Update(task model.Task) (t model.Task, err error) {
	t = task
	err = m.DB.Table(table).Where("id = ?", task.ID).Update("done", task.Done).Error
	return
}
