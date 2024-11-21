package entity

import (
	"github.com/google/uuid"
)

type Post struct {
	ID uuid.UUID
	Author *User
}