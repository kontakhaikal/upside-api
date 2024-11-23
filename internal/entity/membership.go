package entity

import "github.com/google/uuid"

type Role string

var Author Role = "author"

var Admin Role = "admin"

var Member Role = "member"

type Membership struct {
	ID     uuid.UUID `gorm:"column:id;primaryKey"`
	UserID uuid.UUID `gorm:"column:user_id"`
	SideID uuid.UUID `gorm:"column:side_id"`
	Role   Role      `gorm:"column:role"`
}
