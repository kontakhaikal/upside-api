package entity

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID uuid.UUID			`gorm:"column:id;primaryKey"`
	Body string				`gorm:"column:body"`
	CreatedAt time.Time		`gorm:"column:created_at"`
	LastUpdated time.Time	`gorm:"column:last_updated"`
	AuthorID uuid.UUID		`gorm:"column:author_id"`
	SideID uuid.UUID		`gorm:"column:side_id"`
}