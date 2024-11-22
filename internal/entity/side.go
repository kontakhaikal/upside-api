package entity

import "github.com/google/uuid"

type Side struct {
	ID uuid.UUID	`gorm:"column:id;primaryKey"`
	Name string	`gorm:"column:name"`
}