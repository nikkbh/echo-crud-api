package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Task   string `json:"task"`
	Status string `json:"status"`
}
