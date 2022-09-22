package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Firstname string `gotm:"not null" json:"first_name"`
	Lastname  string `gotm:"not null" json:"last_name"`
	Email     string `gotm:"not null;unique_index" json:"email"`
	Tasks     []Task `json:"tasks"`
}
