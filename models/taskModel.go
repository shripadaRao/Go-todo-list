package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Task struct {
	gorm.Model
	// ID		string			`gorm:"PRIMARY_KEY; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	ID		string		`gorm:"PRIMARY_KEY"`
	User	string		`json:"user"`
	Project	string		`json:"project"`
	Title 	string 		`json:"title"`
	Body 	string 		`json:"body"`
	Status 	string 		`json:"status"`
}

func (task *Task) BeforeCreate(tx *gorm.DB) (err error) {
	task.ID = uuid.NewString()
	return
  }