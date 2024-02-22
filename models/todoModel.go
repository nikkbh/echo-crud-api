package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Task   string `json:"task" gorm:"not null"`
	Status string `json:"status" gorm:"default:'pending'"`
}
